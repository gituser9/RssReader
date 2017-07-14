function RssModalController ($scope, $mdDialog, rssService, modalData, mainService){
    $scope.vm = this;
    $scope.feedUrl = "";
    $scope.modalData = modalData;

    $scope.updateFeedName = function() {
        rssService.setNewFeedName($scope.modalData.Feed.Id, $scope.modalData.Feed.Name);
        $scope.cancel();
    };

    $scope.hide = function() {
        $mdDialog.hide();
    };

    $scope.cancel = function() {
        $mdDialog.cancel();
    };

    $scope.addFeed = function() {
        if (!$scope.feedUrl || !$scope.feedUrl.trim().length) {
            return;
        }
        rssService.addFeed($scope.feedUrl, mainService.currentUserId);
        $scope.cancel();
    };

    $scope.delete = function() {
        rssService.delete($scope.modalData.Feed.Id);
        $scope.cancel();
    };

    $scope.toggleUnread = function() {
        rssService.setUnread($scope.modalData.Settings.UnreadOnly);
        $scope.cancel();
    };
}
RssModalController.$inject = [
    "$scope",
    "$mdDialog",
    "rssService",
    "modalData",
    "mainService"
];