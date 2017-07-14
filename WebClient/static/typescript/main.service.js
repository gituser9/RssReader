function MainService ($http, $mdDialog) {

    var factory = {
        modalUrl: "static/html/modals/"
    };

    factory.getSettings = function (userId) {
        var config = {};
        config.params = { id: userId };

        return $http.get("/get-settings", config).then(function(response) {
            return response.data;
        });
    };

    factory.updateSettings = function (userId) {
        var config = {};
        config.params = { id: userId };

        $http.get("/get-settings", config).then(function(response) {
            factory.settings = response.data;
            var storage = window.localStorage;
            var userStr = storage.getItem("RssReaderUser");
            var user = JSON.parse(userStr);
            user.Settings = response.data;
            storage.setItem("RssReaderUser", JSON.stringify(user));
        });
    };

    factory.auth = function (username, password) {
        return $http.post('/auth', { username: username, password: password });
    };

    factory.registration = function (username, password) {
        return $http.post('/registration', { username: username, password: password });
    };

    /*factory.openAuthModal = function () {
        $mdDialog.show({
            controller: ModalController,
            templateUrl: factory.modalUrl + "authModal.html",
            parent: angular.element(document.body),
            clickOutsideToClose: false,
            locals: {
                modalData: null
            }
        });
    };*/

    factory.openModal = function (template, ctrl, modalData) {
        return $mdDialog.show({
            controller: ctrl,
            templateUrl: factory.modalUrl + template,
            parent: angular.element(document.body),
            clickOutsideToClose: true,
            locals: {
                modalData: angular.copy(modalData)
            }
        });
    };

    factory.setSettings = function (settings) {
        $http.post('/set-settings', settings);
    };

    return factory;
}
MainService.$inject = ["$http", "$mdDialog"];
