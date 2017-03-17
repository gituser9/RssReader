function VkController ($scope, $timeout, vkService, mainService) {

    $scope.$watch(function() {
        $scope.model = vkService.model;
    });

    $scope.getVkNews = function() {
        vkService.getVkNews($scope.userId);
    };

    $scope.getPageData = function () {
        vkService.getPageData($scope.userId);
    };
}
VkController.$inject = [
    "$scope",
    "$timeout",
    "vkService",
    "mainService"
];