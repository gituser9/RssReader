
angular.module("app", ["ngSanitize", "bw.paging", "ngFileUpload", "ngMaterial", 'infinite-scroll'])
    .controller("mainCtrl", MainController)
    .controller("vkCtrl", VkController)
    .controller("rssCtrl", RssController)
    .controller("twitterCtrl", TwitterCtrl)
    .service("mainService", MainService)
    .service("vkService", VkService)
    .service("rssService", RssService)
    .service("twitterService", TwitterService)

    /*.config(["$mdThemingProvider", function ($mdThemingProvider) {
        $mdThemingProvider.theme('default')
            .primaryPalette('teal')
            .accentPalette('blue');
    }])*/;
