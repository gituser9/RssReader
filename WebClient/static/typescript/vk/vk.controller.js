function VkController ($scope, $timeout, vkService, mainService) {

    $scope.searchVkGroup = 0;
    $scope.filters = {
        GroupId: 0,
        Page: 1
    };

    $scope.$watch(function() {
        $scope.model = vkService.model;
    });

    $scope.getVkNews = function () {
        ++$scope.filters.Page;  // for scroll
        vkService.getVkNews($scope.userId, $scope.filters.Page);
    };

    $scope.getPageData = function () {
        vkService.getPageData($scope.userId);
    };

    $scope.loadComments = function (news) {
        vkService.loadComments(news);
    };

    $scope.getByFilters = function () {
        vkService.getByFilters($scope.filters);
    };

    $scope.test = function () {
        console.log('DDD');
    }
}
VkController.$inject = [
    "$scope",
    "$timeout",
    "vkService",
    "mainService"
];