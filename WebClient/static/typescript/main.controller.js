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
            this.$scope.userId = mainService.currentUserId;
            this.$scope.settings = mainService.settings;
        });
    }

    setStartedSource(settings) {
        if (settings.RssEnabled) {
            this.rssService.getAll(settings.UserId);
            return;
        }

        if (settings.VkNewsEnabled) {
            this.vkService.getPageData(settings.UserId);
            return;
        }

        if (settings.TwitterEnabled) {
            this.twitterService.getPageData(settings.UserId);
            return;
        }
    };

    init() {
        let storage = window.localStorage;
        let userStr = storage.getItem("RssReaderUser");

        if (userStr) {
            let user = JSON.parse(userStr);

            this.mainService.updateSettings(user.Id);
            this.setStartedSource(user.Settings);

            this.mainService.currentUserId = user.Id;
            this.$scope.isAuth = true;
            this.$scope.username = user.Name;
        } else {
            // modal for auth
            this.mainService.openModal("authModal.html", ModalController, null);
        }
    };

    logout() {
        let storage = window.localStorage;
        storage.removeItem("RssReaderUser");

        // emit event?
        this.mainService.openModal("authModal.html", ModalController, null);
    };

    showRss() {
        this.$scope.currentSource = this.$scope.Sources.Rss;

        if (!this.rssService.feeds || this.rssService.feeds.length === 0) {
            this.rssService.getAll(this.$scope.userId);
        }
    };

    showVk() {
        this.$scope.currentSource = this.$scope.Sources.Vk;

        if (this.vkService.model.VkNews.length === 0) {
            this.vkService.getPageData(this.$scope.userId);
        }
    };

    showTwitter() {
        this.$scope.currentSource = this.$scope.Sources.Twitter;

        if (this.twitterService.model.News.length === 0) {
            this.twitterService.getPageData(this.$scope.userId);
        }
    };

    openMenu($mdOpenMenu, ev) {
        $mdOpenMenu(ev);
    };

    openSettings() {
        this.mainService.getSettings(this.$scope.userId).then((response) => {
            let modalData = { Settings: response };
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