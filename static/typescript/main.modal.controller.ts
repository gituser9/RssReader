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
        Rss: Feeds;
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
                console.log(modalData);
            }
        }

        public updateFeedName(): void {
            this.mainService.setNewFeedName(this.modalData.Rss.Id, this.$scope.modalData.Rss.Name);
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
            this.mainService.delete(this.modalData.Rss.Id);
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
                    let storage = window.localStorage;
                    storage.setItem("RssReaderUser", JSON.stringify(data.User));
                }
            });
        }
    }
}