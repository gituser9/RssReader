class MainService {
    constructor($http, $mdDialog, utilService) {
        this.$http = $http;
        this.$mdDialog = $mdDialog;
        this.utilService = utilService
        this.modalUrl = "static/html/modals/";
        this.settings = {};
    }

    getSettings() {
        return this.utilService.httpGet("/users/settings", (data) => {
            this.settings = data;
            return data;
        });
    };

    getSettingsPromise() {
        return this.$http.get("/users/settings")
    };

    updateSettings() {
        this.utilService.httpGet("/users/settings", (data) => {
            this.settings = data;
            let storage = window.localStorage;
            storage.setItem('settings', JSON.stringify(this.settings))
            // let userStr = storage.getItem("RssReaderUser");
            // let user = JSON.parse(userStr);
            // user.Settings = response.data;
        });
    };

    auth(username, password) {
        return this.$http.post('/auth', { username: username, password: password });
    };

    registration(username, password) {
        return this.$http.post('/registration', { username: username, password: password });
    };

    openModal(template, ctrl, modalData) {
        return this.$mdDialog.show({
            controller: ctrl,
            templateUrl: this.modalUrl + template,
            parent: angular.element(document.body),
            clickOutsideToClose: true,
            locals: {
                modalData: angular.copy(modalData)
            }
        });
    };

    setSettings(settings) {
        this.settings = settings
        this.utilService.httpPut('/users/settings', settings);
    };
}
MainService.$inject = ["$http", "$mdDialog", "utilService"];

angular.module('app').service('mainService', MainService);