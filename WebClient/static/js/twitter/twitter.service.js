class TwitterService {
    constructor(utilService) {
        this.utilService = utilService;
        this.model = {
            News: [],
            Sources: [],
            IsLoad: false,
            IsAll: false,
            SourceMap: {},
            IsSimpleVersion: false
        };
    }

    getPageData() {
        this.model.IsLoad = true;

        this.utilService.httpGet("/twitter", (data) => {
            this.model.News = data.News;
            this.model.Sources = data.Sources;
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
        this.utilService.httpGet(`/twitter/news?page=${page}`, (data) => {
            this.model.IsLoad = false;

            if (data.length === 0) {
                this.model.IsAll = true;
                return;
            }
            for (let item of data) {
                this.model.News.push(item);
            }
        });
    };

    getSources() {
        this.utilService.httpGet("/twitter/sources", (data) => {
            this.model.Sources = data;
        });
    };

    getByFilters(filters) {
        this.model.News = [];
        let requestUrl = `/twitter/search'?search_string=${filters.SearchString}&source_id=${filters.SourceId}`

        this.utilService.httpGet(requestUrl, (data) => {
            this.model.News = data;
        });
    };

    search(searchString, sourceId) {
        this.model.IsSearch = true;
        let requestUrl = `/twitter/search'?search_string=${searchString}&source_id=${sourceId}`

        this.utilService.httpGet(requestUrl, (data) => {
            this.model.News = data;
        });
    };
}
TwitterService.$inject = ['utilService'];

angular.module('app').service('twitterService', TwitterService);