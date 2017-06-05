function MainController ($scope, mainService, vkService, rssService, twitterService) {

    var setStartedSource = function(settings) {
        if (settings.RssEnabled) {
            rssService.getAll(settings.UserId);
            return;
        }

        if (settings.VkNewsEnabled) {
            vkService.getPageData(settings.UserId);
            return;
        }

        if (settings.TwitterEnabled) {
            twitterService.getPageData(settings.UserId);
            return;
        }
    };

    $scope.Sources = {
        Rss: 1,
        Vk: 2,
        Twitter: 3
    };
    $scope.currentSource = $scope.Sources.Rss;
    $scope.isAuth = false;


    $scope.$watch(function () {
        $scope.userId = mainService.currentUserId;
        $scope.settings = mainService.settings;
    });

    $scope.init = function () {
        var storage = window.localStorage;
        var userStr = storage.getItem("RssReaderUser");

        if (userStr) {
            var user = JSON.parse(userStr);

            // mainService.settings = user.Settings;
            mainService.updateSettings(user.Id);
            setStartedSource(user.Settings);

            mainService.currentUserId = user.Id;
            $scope.isAuth = true;
            $scope.username = user.Name;
        } else {
            // modal for auth
            // mainService.openAuthModal();
            mainService.openModal("authModal.html", ModalController, null);
        }
    };

    $scope.logout = function () {
        var storage = window.localStorage;
        storage.removeItem("RssReaderUser");

        // emit event?
        mainService.openAuthModal();
    };

    $scope.showRss = function () {
        $scope.currentSource = $scope.Sources.Rss;

        if (!rssService.feeds || rssService.feeds.length === 0) {
            rssService.getAll($scope.userId);
        }
    };

    $scope.showVk = function () {
        $scope.currentSource = $scope.Sources.Vk;

        if (vkService.model.VkNews.length === 0) {
            vkService.getPageData($scope.userId);
        }
    };

    $scope.showTwitter = function () {
        $scope.currentSource = $scope.Sources.Twitter;

        if (twitterService.model.news.length === 0) {
            twitterService.getPageData($scope.userId);
        }
    };

    $scope.openMenu = function($mdOpenMenu, ev) {
        $mdOpenMenu(ev);
    };

/*
Modals
================================================================================
*/

    $scope.openSettings = function() {
        var storage = window.localStorage;
        var userStr = storage.getItem("RssReaderUser");
        var user = JSON.parse(userStr);

        mainService.getSettings(user.Id).then(function (response) {
            var modalData = {};
            modalData.Settings = response;
            mainService.openModal("settingModal.html", ModalController, modalData);
        });
    };

/*
Private
================================================================================ */



}
MainController.$inject = [
    "$scope",
    "mainService",
    "vkService",
    "rssService",
    "twitterService"
];
