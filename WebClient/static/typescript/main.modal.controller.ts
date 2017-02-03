/// <reference path="_all.ts" />

module main {
    "use strict";

    import IDialogService = angular.material.IDialogService;

    interface IModalScope extends ng.IScope {
        vm: ModalController;
        modalData: ModalData;
        feedUrl: string;
        username: string;
        password: string;
        errorMessage: string;
    }

    export class ModalData {
        Feed: Feeds;
        Settings: Settings;
    }

    export class ModalController {
        public static $inject = [
            "$scope",
            "$mdDialog",
            "mainService",
            "modalData"
        ];

        constructor(
            private $scope: IModalScope,
            private $mdDialog: IDialogService,
            private mainService: MainService,
            private modalData?: ModalData
        ) {
            $scope.vm = this;
            $scope.feedUrl = "";

            if (modalData != null) {
                $scope.modalData = modalData;
            }
        }

        public updateFeedName(): void {
            this.mainService.setNewFeedName(this.modalData.Feed.Id, this.$scope.modalData.Feed.Name);
            this.cancel();
        }

        public hide(): void {
            this.$mdDialog.hide();
        }

        public cancel(): void {
            this.$mdDialog.cancel();
        }

        public addFeed(): void {
            if (!this.$scope.feedUrl || !this.$scope.feedUrl.trim().length) {
                return;
            }

            this.mainService.addFeed(this.$scope.feedUrl);
            this.cancel();
        }

        public delete(): void {
            this.mainService.delete(this.modalData.Feed.Id);
            this.cancel();
        }

        public toggleUnread(): void {
            this.mainService.setUnread(this.$scope.modalData.Settings.UnreadOnly);
            this.cancel();
        }

        public auth(): void {
            this.mainService.auth(this.$scope.username, this.$scope.password).then((response: ng.IHttpPromiseCallbackArg<User>) => {
                if (response.data != null) {
                    this.cancel();
                    this.mainService.currentUserId = (<User> response.data).Id;
                    this.mainService.settings = response.data.Settings;
                    this.mainService.getAll(this.mainService.currentUserId);

                    let storage = window.localStorage;
                    storage.setItem("RssReaderUser", JSON.stringify(response.data));

                }
            });
        }

        public registration(): void {
            this.mainService.registration(this.$scope.username, this.$scope.password).then((response: ng.IHttpPromiseCallbackArg<RegistrationData>) => {
                if (response.data != null) {
                    let data = <RegistrationData> response.data;

                    if (data.User == null) {
                        this.$scope.errorMessage = data.Message;
                        return
                    }

                    this.cancel();

                    this.$scope.errorMessage = "";
                    this.mainService.currentUserId = data.User.Id;
                    this.mainService.settings = data.User.Settings;
                    
                    let storage = window.localStorage;
                    storage.setItem("RssReaderUser", JSON.stringify(data.User));
                }
            });
        }

        public saveSettings(settings: Settings): void {
            this.mainService.setSettings(JSON.stringify(settings));
            this.mainService.settings = settings;
            
            let storage = window.localStorage;
            let userStr = storage.getItem("RssReaderUser");
            let user = <User> JSON.parse(userStr);
            user.Settings = settings;

            storage.setItem("RssReaderUser", JSON.stringify(user));
            this.cancel();
        }
    }
}
