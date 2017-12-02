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

        if (userStr) {
            let user = JSON.parse(userStr);
            this.userId = user.Settings.UserId;
        }
    }

    getPageData() {
        this.model.IsLoad = true;
        let config = {};
        config.params = { id: this.userId };

        this.$http.get("/get-twitter-page", config).then((response) => {
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
        config.params = { id: this.userId, page: page };

        this.$http.get("/get-twitter-news", config).then((response) => {
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
        let config = {};
        config.params = { id: this.userId };

        this.$http.get("/get-", config).then((response) => {
            this.model.Sources = response.data;
        });
    };

    getByFilters(filters) {
        this.model.News = [];
        let data = {
            SearchString: filters.SearchString,
            SourceId: Number(filters.SourceId)
        };
        this.$http.post('/get-twitter-news-by-filters', data).then((response) => {
            this.model.News = response.data;
        });
    };

    search(searchString, sourceId) {
        this.model.IsSearch = true;
        let data = {
            SearchString: searchString,
            SourceId: sourceId
        };
        this.$http.post('/search-twitter-news', data).then((response) => {
            this.model.News = response.data;
        });
    };
}
TwitterService.$inject = ['$http'];

angular.module('app').service('twitterService', TwitterService);