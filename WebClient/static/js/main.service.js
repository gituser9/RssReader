class MainService {
    constructor($http, $mdDialog) {
        this.$http = $http;
        this.$mdDialog = $mdDialog;
        this.modalUrl = "static/html/modals/";
        this.settings = {};
    }

    getSettings() {
        return this.$http.get("/users/settings").then((response) => {
            this.settings = response.data;
            return response.data;
        });
    };

    updateSettings() {
        this.$http.get("/users/settings").then((response) => {
            this.settings = response.data;
            // let storage = window.localStorage;
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
        this.$http.post('/users/settings', settings);
    };
}
MainService.$inject = ["$http", "$mdDialog"];

angular.module('app').factory('myInterceptor', ['$log', function($log) {
    return {
        'response': function(response) {
           // same as above
           console.log("RESPONSE");
           console.log(response.status);
           console.log(response);

           if (response.status == 403) {
                let rtoken = localStorage.getItem('rtoken')

                if (rtoken === null || rtoken === '') {
                    return response;
                }

                localStorage.setItem('rtoken', '')
                let xhr = new XMLHttpRequest();
                xhr.open('PUT', '/users/refresh', false);
                xhr.send({ 'token': rtoken });

                if (xhr.status != 200) {
                    return response;
                } else {
                    let data = JSON.parse(xhr.responseText)
                    localStorage.setItem('rtoken', data.refresh_token)
                    localStorage.setItem('token', data.token)
                }
           }

           return response;
        },
        'request': function(config) {
            config.headers['Authorization'] = 'Bearer ' localStorage.getItem('token')
            // console.log('REQUEST');
            console.log(config);
            
            return config;
        },
      };
}]);
angular.module('app').config(['$httpProvider', function($httpProvider) {
    $httpProvider.interceptors.push('myInterceptor');
}]);
angular.module('app').service('mainService', MainService);