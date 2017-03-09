import IDialogService = angular.material.IDialogService;

import { Settings, User, RegistrationData, ModalData } from './models/generalModels';
import { ModalController } from './main.modal.controller';


export class MainService {
    public currentUserId: number;
    public settings: Settings;

    public static $inject = ["$http", "$mdDialog"];

    constructor(private $http: ng.IHttpService, private $mdDialog: IDialogService) {
        this.$http = $http;
    }

    public getSettings(userId: number): ng.IPromise<Settings> {
        let config: ng.IRequestShortcutConfig = {};
        config.params = { id: userId };

        return this.$http.get("/get-settings", config).then((response: ng.IHttpPromiseCallbackArg<Settings>): Settings => {
            return <Settings> response.data;
        });
    }

    public updateSettings(userId: number): void {
        let config: ng.IRequestShortcutConfig = {};
        config.params = { id: userId };

        this.$http.get("/get-settings", config).then((response: ng.IHttpPromiseCallbackArg<Settings>): void => {
            this.settings = <Settings> response.data;
        });
    }

    

    public auth(username: string, password: string): ng.IHttpPromise<User> {
        return this.$http.post('/auth', { username: username, password: password });
    }

    public registration(username: string, password: string): ng.IHttpPromise<RegistrationData> {
        return this.$http.post('/registration', { username: username, password: password });
    }

    public openAuthModal(): void {
        this.$mdDialog.show({
            controller: ModalController,
            templateUrl: "static/html/modals/authModal.html",
            parent: angular.element(document.body),
            clickOutsideToClose: false,
            locals: {
                modalData: null
            }
        });
    }

    // todo: ctrl type
    public openModal(url: string, ctrl: any, modalData?: ModalData): void {
        this.$mdDialog.show({
            controller: ctrl,
            templateUrl: url,
            parent: angular.element(document.body),
            clickOutsideToClose: true,
            locals: {
                modalData: angular.copy(modalData)
            }
        });
    }

    public setSettings(settings: Settings): void {
        this.$http.post('/set-settings', settings);
    }

}
