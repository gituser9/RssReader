class MainController {
    constructor($scope, mainService, vkService, rssService, twitterService) {
        this.$scope = $scope;
        this.mainService = mainService;
        this.vkService = vkService;
        this.rssService = rssService;
        this.twitterService = twitterService;

        this.$scope.Sources = {
            Rss: 1,
            Vk: 2,
            Twitter: 3
        };
        this.$scope.currentSource = $scope.Sources.Rss;
        this.$scope.isAuth = false;

        this.$scope.$watch(() => {
            //  = mainService.currentUserId;
            // this.$scope.settings = mainService.settings;
        });
    }

    setStartedSource(settings) {
        // if (settings.RssEnabled) {
        if (true) {
            this.rssService.getAll();
            return;
        }

        if (settings.VkNewsEnabled) {
            this.vkService.getPageData();
            return;
        }

        if (settings.TwitterEnabled) {
            this.twitterService.getPageData();
            return;
        }
    };

    init() {
        let storage = window.localStorage;
        let token = storage.getItem("token");

        if (token) {
            this.mainService.updateSettings();
            this.setStartedSource(user.Settings);

            this.$scope.isAuth = true;
        } else {
            // modal for auth
            this.mainService.openModal("authModal.html", ModalController, null);
        }
    };

    logout() {
        let storage = window.localStorage;
        storage.removeItem("token");
        storage.removeItem("rtoken");

        // emit event?
        this.mainService.openModal("authModal.html", ModalController, null);
    };

    showRss() {
        this.$scope.currentSource = this.$scope.Sources.Rss;

        if (!this.rssService.feeds || this.rssService.feeds.length === 0) {
            this.rssService.getAll();
        }
    };

    showVk() {
        this.$scope.currentSource = this.$scope.Sources.Vk;

        if (this.vkService.model.VkNews.length === 0) {
            this.vkService.getPageData();
        }
    };

    showTwitter() {
        this.$scope.currentSource = this.$scope.Sources.Twitter;

        if (this.twitterService.model.News.length === 0) {
            this.twitterService.getPageData();
        }
    };

    openMenu($mdOpenMenu, ev) {
        $mdOpenMenu(ev);
    };

    openSettings() {
        this.mainService.getSettings().then((response) => {
            let modalData = { Settings: response.data };
            this.mainService.openModal("settingModal.html", ModalController, modalData);
        });
    }
}
MainController.$inject = [
    "$scope",
    "mainService",
    "vkService",
    "rssService",
    "twitterService"
];


angular.module('app').controller('mainCtrl', MainController);