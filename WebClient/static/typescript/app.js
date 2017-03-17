
angular.module("app", ["ngSanitize", "ui.bootstrap", "ngFileUpload", "ngMaterial"])
    .controller("mainCtrl", MainController)
    // .controller("modalCtrl", ModalController)
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
