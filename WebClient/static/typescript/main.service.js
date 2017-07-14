class MainService {
    constructor($http, $mdDialog) {
        this.$http = $http;
        this.$mdDialog = $mdDialog;
        this.modalUrl = "static/html/modals/";
        this.settings = {};
    }

    getSettings(userId) {
        let config = {};
        config.params = { id: userId };

        return this.$http.get("/get-settings", config).then((response) => {
            return response.data;
        });
    };

    updateSettings(userId) {
        let config = {};
        config.params = { id: userId };

        this.$http.get("/get-settings", config).then((response) => {
            this.settings = response.data;
            let storage = window.localStorage;
            let userStr = storage.getItem("RssReaderUser");
            let user = JSON.parse(userStr);
            user.Settings = response.data;
            storage.setItem("RssReaderUser", JSON.stringify(user));
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
        this.$http.post('/set-settings', settings);
    };
}
MainService.$inject = ["$http", "$mdDialog"];

angular.module('app').service('mainService', MainService);