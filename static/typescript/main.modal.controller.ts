/// <reference path="_all.ts" />

module main {
    "use strict";

    import IDialogService = angular.material.IDialogService;

    interface IModalScope extends ng.IScope {
        vm: ModalController;
        modalData: ModalData;
        feedUrl: string;
    }

    export class ModalData {
        Rss: Rss;
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
            this.mainService.setNewFeedName(this.modalData.Rss.ID, this.$scope.modalData.Rss.RssName);
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
            this.mainService.delete(this.modalData.Rss.ID);
            this.cancel();
        }

        public toggleUnread(): void {
            this.mainService.setUnread(this.$scope.modalData.Settings.UnreadOnly);
            this.cancel();
        }
    }
}