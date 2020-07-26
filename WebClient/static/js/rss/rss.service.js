class RssService {
    constructor($http, utilService) {
        this.articles = [];
        this.articlesCount = 0;
        this.showWaitBar = false;

        this.http = $http;
        this.utilService = utilService
    }

    getArticles(feedId, page) {
        this.utilService.httpGet(`/rss/${feedId}/articles?page=${page}`, (data) => {
            this.articles = data.Articles;
            this.articlesCount = data.Count;
            this.showArticle = false;
        });
    };

    search(filters) {
        let config = {};
        config.params = {
            searchString: filters.searchText,
            isBookmark: filters.searchInBookmark,
            feedId: filters.searchFeed
        };

        let url = `/rss/search?search_string=${filters.searchText}&feed_id=${filters.searchFeed}&is_bookmark=${filters.searchInBookmark}`
        this.utilService.httpGet(url, (data) => {
            this.articles = data.Articles;
            this.articlesCount = data.Count;
            this.showArticle = false;
        });
    };

    getArticle(article, feedId) {
        console.log(article)
        this.utilService.httpGet(`/rss/${feedId}/articles/${article.Id}`, (data) => {
            this.article = data;
            this.showArticle = true;

            if (this.article) {
                this.articles.forEach((item) => {
                    if (item.Id === article.Id) {
                        item.IsRead = true;
                    }
                });
            }
        });
    };

    getArticlePromise(article, callback) {
        this.utilService.httpGet(`/rss/${article.FeedId}/articles/${article.Id}`, (data) => {
            callback(data)
        });
    };

    getAll() {
        this.utilService.httpGet('/rss', (data) => {
            this.feeds = data;
            this.showWaitBar = false;
        })
    };

    addFeed(url) {
        this.utilService.httpPost("/rss", { url: url }, (data) => {
            this.feeds = data;
        });
    };

    deleteFeed(id) {
        this.utilService.httpDelete(`/rss/${id}`, (data) => {
            this.feeds = data;
        });
    };

    setNewFeedName(id, name) {
        this.utilService.httpPut(`/rss/${id}`, { name: name }, () => {
            this.getAll()
        });
    };

    updateAll() {
        this.showWaitbar = true;

        this.utilService.httpGet('/rss', (data) => {
            this.feeds = data;
            this.showWaitbar = false;
        });
    };

    toggleBookmark(articleId, page, isBookmark, isBookmarkPage, feedId) {
        this.utilService.httpPut(`/rss/${feedId}/articles/${articleId}`, { isBookmark: isBookmark, page: page }, (data) => {
            if (!data) {
                return;
            }
            if (isBookmarkPage) {
                this.getBookmarks(page);
            } else {
                this.getArticles(feedId, page);
            }
        });
    };

    getBookmarks(page) {
        let url = `/rss/articles/bookmarks?page=${page}`
        this.utilService.httpGet(url, (data) => {
            this.articles = data.Articles;
            this.articlesCount = data.Count;
            this.showArticle = false;
        });
    };

    markReadAll(feedId) {
        let config = { is_read_all: true };

        this.utilService.httpPut(`/rss/${feedId}`, config, (data) => {
            this.articles = data.Articles;
            this.articlesCount = data.Count;
        });
    };

    createOpml(callback) {
        this.utilService.httpGet('/rss/opml', (data) => {
            callback(data)
        });
    };

    markAsRead(id, feedId, isRead) {
        let params = { article_id: id, is_read: isRead };

        this.utilService.httpPut(`/rss/${feedId}/articles/${id}`, params, null);
    };

    setUnread(id, feedId) {
        this.utilService.httpPut(`/rss/${feedId}/articles/${id}`, { is_read: false, article_id: id }, null);
    };
}
RssService.$inject = ['$http', 'utilService'];

angular.module('app').service('rssService', RssService);