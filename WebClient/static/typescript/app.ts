/// <reference path="_all.ts" />

module main {
    "use strict";

    angular.module("app", ["ngSanitize", "ui.bootstrap", "ngFileUpload", "ngMaterial"])
        .controller("mainCtrl", MainController)
        .controller("modalCtrl", ModalController)
        .service("mainService", MainService)
        .config(["$mdThemingProvider", ($mdThemingProvider: angular.material.IThemingProvider) => {
            $mdThemingProvider.theme('default')
                .primaryPalette('teal')
                .accentPalette('blue');
        }]);
}
