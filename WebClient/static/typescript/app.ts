/// <reference path="_all.ts" />
import { MainController } from './main.controller';
import { ModalController } from './main.modal.controller';
import { MainService } from './main.service';
import { VkController } from './vk/vk.controller';
import { VkService } from './vk/vk.service';
import { RssController } from './rss/rss.controller';
import { RssService } from './rss/rss.service';


module main {
    "use strict";

    angular.module("app", ["ngSanitize", "ui.bootstrap", "ngFileUpload", "ngMaterial"])
        .controller("mainCtrl", MainController)
        // .controller("modalCtrl", ModalController)
        .controller("vkCtrl", VkController)
        .controller("rssCtrl", RssController)
        .service("mainService", MainService)
        .service("vkService", VkService)
        .service("rssService", RssService)
        .config(["$mdThemingProvider", ($mdThemingProvider: angular.material.IThemingProvider) => {
            $mdThemingProvider.theme('default')
                .primaryPalette('teal')
                .accentPalette('blue');
        }]);
}
