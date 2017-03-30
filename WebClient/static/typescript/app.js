
angular.module("app", ["ngSanitize", "ui.bootstrap", "ngFileUpload", "ngMaterial", 'infinite-scroll'])
    .controller("mainCtrl", MainController)
    .controller("vkCtrl", VkController)
    .controller("rssCtrl", RssController)
    .service("mainService", MainService)
    .service("vkService", VkService)
    .service("rssService", RssService)
    /*.config(["$mdThemingProvider", function ($mdThemingProvider) {
        $mdThemingProvider.theme('default')
            .primaryPalette('teal')
            .accentPalette('blue');
    }])*/;
