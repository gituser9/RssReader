/// <reference path="../_all.ts" />

import IDialogService = angular.material.IDialogService;

import { Feeds } from './rss.models';
import { ModalData } from '../models/generalModels';
import { RssService } from './rss.service';


interface IModalScope extends ng.IScope {
    vm: RssModalController;
    modalData: ModalData;
    feedUrl: string;
}

export class RssModalController {
    public static $inject = [
        "$scope",
        "$mdDialog",
        "rssService",
        "modalData"
    ];

    constructor(
        private $scope: IModalScope,
        private $mdDialog: IDialogService,
        private rssService: RssService,
        private modalData?: ModalData
    ) {
        $scope.vm = this;
        $scope.feedUrl = "";

        if (modalData != null) {
            $scope.modalData = modalData;
        }
    }

    public updateFeedName(): void {
        this.rssService.setNewFeedName(this.modalData.Feed.Id, this.$scope.modalData.Feed.Name);
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

        this.rssService.addFeed(this.$scope.feedUrl);
        this.cancel();
    }

    public delete(): void {
        this.rssService.delete(this.modalData.Feed.Id);
        this.cancel();
    }

    public toggleUnread(): void {
        this.rssService.setUnread(this.$scope.modalData.Settings.UnreadOnly);
        this.cancel();
    }
}