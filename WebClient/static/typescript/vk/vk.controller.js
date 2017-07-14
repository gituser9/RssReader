function VkController ($scope, $timeout, vkService, mainService) {

    $scope.searchVkGroup = 0;
    $scope.filters = {
        GroupId: 0,
        Page: 1,
        SearchString: ''
    };

    $scope.$watch(function() {
        $scope.model = vkService.model;
    });

    $scope.refresh = function () {
        $scope.filters.Page = 0;
        $scope.goToTop();
        vkService.model.IsSearch = false;
        vkService.model.VkNews = [];
        $scope.getVkNews();
    };

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

    $scope.search = function () {
        vkService.search($scope.filters.SearchString, $scope.filters.GroupId);
    };
}
VkController.$inject = [
    "$scope",
    "$timeout",
    "vkService",
    "mainService"
];