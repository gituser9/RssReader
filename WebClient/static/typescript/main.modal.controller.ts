/// <reference path="_all.ts" />

import { Settings, User, RegistrationData, ModalData } from './models/generalModels';
import { MainService } from './main.service';

    
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

    public hide(): void {
        this.$mdDialog.hide();
    }

    public cancel(): void {
        this.$mdDialog.cancel();
    }

    

    public auth(): void {
        this.mainService.auth(this.$scope.username, this.$scope.password).then((response: ng.IHttpPromiseCallbackArg<User>) => {
            if (response.data != null) {
                this.cancel();
                this.mainService.currentUserId = (<User> response.data).Id;
                this.mainService.settings = response.data.Settings;
                // this.mainService.getAll(this.mainService.currentUserId);

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

    public saveSettings(): void {
        this.mainService.setSettings(this.modalData.Settings);
        this.mainService.settings = this.modalData.Settings;
        
        let storage = window.localStorage;
        let userStr = storage.getItem("RssReaderUser");
        let user = <User> JSON.parse(userStr);
        user.Settings = this.modalData.Settings;

        storage.setItem("RssReaderUser", JSON.stringify(user));
        this.cancel();
    }
}
