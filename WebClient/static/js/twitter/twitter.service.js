class TwitterService {
    constructor($http) {
        this.$http = $http;
        this.model = {
            News: [],
            Sources: [],
            IsLoad: false,
            IsAll: false,
            SourceMap: {},
            IsSimpleVersion: false
        };
        let storage = window.localStorage;
        let userStr = storage.getItem("RssReaderUser");
    }

    getPageData() {
        this.model.IsLoad = true;

        this.$http.get("/twitter/page").then((response) => {
            this.model.News = response.data.News;
            this.model.Sources = response.data.Sources;
            this.model.IsLoad = false;

            this.model.Sources.forEach((item) => {
                this.model.SourceMap[item.Id] = {
                    image: item.Image,
                    name: item.Name,
                    link: item.Url,
                    screenName: item.ScreenName
                }
            });
        });
    };

    getNews(page) {
        if (this.model.IsLoad || this.model.IsAll) {
            return;
        }

        this.model.IsLoad = true;
        let config = {};
        config.params = { page: page };

        this.$http.get("/twitter/news", config).then((response) => {
            this.model.IsLoad = false;

            if (response.data.length === 0) {
                this.model.IsAll = true;
                return;
            }
            for (let item of response.data) {
                this.model.News.push(item);
            }
        });
    };

    getSources() {
        this.$http.get("/twitter/sources").then((response) => {
            this.model.Sources = response.data;
        });
    };

    getByFilters(filters) {
        this.model.News = [];
        let data = {
            SearchString: filters.SearchString,
            SourceId: Number(filters.SourceId)
        };
        this.$http.get('/twitter/news', data).then((response) => {
            this.model.News = response.data;
        });
    };

    search(searchString, sourceId) {
        this.model.IsSearch = true;
        let data = {
            SearchString: searchString,
            SourceId: sourceId
        };
        this.$http.get('/twitter/news', data).then((response) => {
            this.model.News = response.data;
        });
    };
}
TwitterService.$inject = ['$http'];

angular.module('app').service('twitterService', TwitterService);