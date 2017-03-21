function VkController ($scope, $timeout, vkService, mainService) {

    $scope.searchVkGroup = 0;
    $scope.filters = {
        GroupId: 0
    };

    $scope.$watch(function() {
        $scope.model = vkService.model;
    });

    $scope.getVkNews = function() {
        vkService.getVkNews($scope.userId);
    };

    $scope.getPageData = function () {
        vkService.getPageData($scope.userId);
    };

    $scope.loadComments = function (news) {
        vkService.loadComments(news);
    };

    $scope.getByFilters = function () {
        vkService.getByFilters($scope.filters);
    }
}
VkController.$inject = [
    "$scope",
    "$timeout",
    "vkService",
    "mainService"
];