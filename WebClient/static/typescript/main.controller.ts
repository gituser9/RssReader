import { Settings, User, ModalData, Sources } from './models/generalModels';
import { MainService } from './main.service';
import { ModalController } from './main.modal.controller';


"use strict";

import IDialogService = angular.material.IDialogService;
import IToastService = angular.material.IToastService;


interface IMainScope extends ng.IScope {
    vm: MainController;
    settings: Settings;
    username: string;
    currentUserId: number;
    sources: Sources;
    currentSource: Sources;
}

export class MainController {
    public static $inject = [
        "$scope",
        "$timeout",
        "$mdDialog",
        "$mdToast",
        "Upload",
        "mainService"
    ];
    private isAuth: boolean;
    private userId: number;

    constructor(
        private $scope: IMainScope, 
        private $timeout: ng.ITimeoutService,
        private $mdDialog: IDialogService, 
        private $mdToast: IToastService,
        private $upload: any, 
        private mainService: MainService
    ) {
        $scope.vm = this;
        this.$scope.currentSource = Sources.Rss;
        this.isAuth = false;

        $scope.$watch(() => {
            this.userId = mainService.currentUserId;
            this.$scope.settings = mainService.settings;
        });            

        let storage = window.localStorage;
        let userStr = storage.getItem("RssReaderUser");

        if (userStr != null) {
            let user = <User> JSON.parse(userStr);

            // this.mainService.settings = user.Settings;
            this.mainService.updateSettings(user.Id);
            // this.mainService.getAll(user.Id);

            this.setStartedSource();

            mainService.currentUserId = user.Id;

            this.isAuth = true;
            this.$scope.username = user.Name;

            // this.$timeout(() => { this.mainService.getAll(user.Id) }, 30);  
        } else {
            // modal for auth
            this.mainService.openAuthModal();
        }            
    }

    public logout(): void {
        let storage = window.localStorage;
        storage.removeItem("RssReaderUser");

        // emit event?
        this.mainService.openAuthModal();
    }

    public showRss(): void {
        this.$scope.currentSource = Sources.Rss;

    }

    public showVk(): void {
        this.$scope.currentSource = Sources.Vk;
    }

        
        
/*
Modals
================================================================================
*/

    public openSettings(): void {
        let storage = window.localStorage;
        let userStr = storage.getItem("RssReaderUser");
        let user = <User> JSON.parse(userStr);

        this.mainService.getSettings(user.Id).then((response: Settings): void => {
            let modalData = new ModalData();
            modalData.Settings = response;
            this.mainService.openModal("static/html/modals/settingModal.html", ModalController, modalData);
        });
    }

        

/*
Private
================================================================================ */

    private setStartedSource(): void {

    }

}
