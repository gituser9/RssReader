function TwitterCtrl($scope, twitterService) {
    'use strict';

    $scope.searchSource = 0;
    $scope.filters = {
        SourceId: 0,
        Page: 1,
        SearchString: ''
    };

    $scope.$watch(function() {
        $scope.model = twitterService.model;
    });

    $scope.getNews = function () {
        ++$scope.filters.Page;  // for scroll
        twitterService.getNews($scope.userId, $scope.filters.Page);
    };

    $scope.getPageData = function () {
        twitterService.getPageData($scope.userId);
    };

    $scope.getByFilters = function () {
        twitterService.getByFilters($scope.filters);
    };

    $scope.search = function () {
        twitterService.search($scope.filters.SearchString, $scope.filters.SourceId);
    };
}
TwitterCtrl.$inject = ['$scope', 'twitterService'];

