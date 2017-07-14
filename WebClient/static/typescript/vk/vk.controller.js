class VkController {
    constructor($scope, vkService) {
        this.vkService = vkService;
        this.$scope = $scope;
        this.$scope.searchVkGroup = 0;
        this.$scope.filters = {
            GroupId: 0,
            Page: 1,
            SearchString: ''
        };
        this.$scope.$watch(() => {
            this.$scope.model = this.vkService.model;
        });
    }

    refresh() {
        this.$scope.filters.Page = 0;
        this.vkService.model.IsSearch = false;
        this.vkService.model.VkNews = [];
        this.getVkNews();
    };

    getVkNews() {
        ++this.$scope.filters.Page;  // for scroll
        this.vkService.getVkNews(this.$scope.userId, this.$scope.filters.Page);
    };

    getPageData() {
        this.vkService.getPageData(this.$scope.userId);
    };

    loadComments(news) {
        this.vkService.loadComments(news);
    };

    getByFilters() {
        this.vkService.getByFilters(this.$scope.filters);
    };

    search() {
        this.vkService.search(this.$scope.filters.SearchString, this.$scope.filters.GroupId);
    };
}
VkController.$inject = ['$scope', 'vkService'];

angular.module('app').controller('vkCtrl', VkController);