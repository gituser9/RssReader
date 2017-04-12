function RssModalController ($scope, $mdDialog, rssService, modalData){
    $scope.vm = this;
    $scope.feedUrl = "";

    console.log(modalData);

    if (modalData != null) {
        $scope.modalData = modalData;
    }

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

        rssService.addFeed($scope.feedUrl);
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
    "modalData"
];