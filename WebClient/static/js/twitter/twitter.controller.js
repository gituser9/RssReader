class TwitterCtrl {
    constructor($scope, twitterService, mainService) {
        this.$scope = $scope;
        this.twitterService = twitterService;

        this.$scope.searchSource = 0;
        this.$scope.filters = {
            SourceId: 0,
            Page: 1,
            SearchString: ''
        };
        this.$scope.$watch(() => {
            this.$scope.model = twitterService.model;
            this.$scope.IsSimpleVersion = mainService.settings.TwitterSimpleVersion;
        });
    }

    getNews() {
        ++this.$scope.filters.Page;  // for scroll
        this.twitterService.getNews(this.$scope.filters.Page);
    };

    getPageData() {
        console.log(this.$scope.IsSimpleVersion);

        this.twitterService.getPageData();
    };

    getByFilters() {
        this.twitterService.getByFilters(this.$scope.filters);
    };

    search() {
        this.twitterService.search(this.$scope.filters.SearchString, Number(this.$scope.filters.SourceId));
    };

    refresh() {
        this.$scope.filters.Page = 0;
        this.twitterService.model.IsSearch = false;
        this.twitterService.model.IsAll = false;
        this.twitterService.model.News = [];
        this.getNews();
    };
}
TwitterCtrl.$inject = ['$scope', 'twitterService', 'mainService'];


angular.module('app').controller('twitterCtrl', TwitterCtrl);