import { User } from '../models/generalModels';
import { VkGroup, VkNews } from './vk.models';
import { VkService } from './vk.service';
import { MainService } from '../main.service';


"use strict";


interface IVkScope extends ng.IScope {
    vm: VkController;
    vkNews: VkNews[];
    vkGroups: VkGroup[];
}

export class VkController {
    public static $inject = [
        "$scope",
        "$timeout",
        "$mdDialog",
        "$mdToast",
        "vkService",
        "mainService"
    ];

    private userId: number;

    constructor(
        private $scope: IVkScope, 
        private $timeout: ng.ITimeoutService,
        private vkService: VkService,
        private mainService: MainService
    ) {
        $scope.vm = this;

        $scope.$watch(() => {
            $scope.vkGroups = vkService.vkGroups;
            $scope.vkNews = vkService.vkNews;
        });

        let storage = window.localStorage;
        let userStr = storage.getItem("RssReaderUser");

        if (userStr != null) {
            let user = <User> JSON.parse(userStr);

            // this.mainService.settings = user.Settings;
            this.mainService.updateSettings(user.Id);
            // this.mainService.getAll(user.Id);

            // this.$timeout(() => { this.mainService.getAll(user.Id) }, 30);  
        } else {
            // modal for auth
            this.mainService.openAuthModal();
        }
    }

    public getVkNews(): void {
        this.vkService.getVkNews(this.userId);
    }
}
angular.module("app").controller("vkCtrl", VkController);