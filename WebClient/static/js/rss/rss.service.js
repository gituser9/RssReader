class RssService {
    constructor($http) {
        this.articles = [];
        this.articlesCount = 0;
        this.showWaitBar = false;

        this.http = $http;
    }

    getArticles(feedId, page) {
        let config = {};
        config.params = { "page": page };

        this.http.get(`/rss/${feedId}/articles`, config).then((response) => {
            this.articles = response.data.Articles;
            this.articlesCount = response.data.Count;
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

        this.http.get("/rss/search", config).then((response) => {
            this.articles = response.data.Articles;
            this.articlesCount = response.data.Count;
            this.showArticle = false;
        });
    };

    getArticle(article) {
        this.http.get(`/rss/${article.FeedId}/articles/${article.Id}`).then((response) => {
            this.article = response.data;
            this.showArticle = true;

            if (this.article) {
                this.articles.forEach(function(item) {
                    if (item.Id === this.article.Id) {
                        item.IsRead = true;
                    }
                });
            }

        });
    };

    getArticlePromise(article) {
        return this.http.get(`/rss/${article.FeedId}/articles/${article.Id}`);
    };

    getAll() {
        this.http.get("/rss").then((response) => {
            this.feeds = response.data;
            this.showWaitBar = false;
        });
    };

    addFeed(url, userId) {
        this.http.post("/rss", { url: url, userId: userId }).then((response) => {
            this.feeds = response.data;
        });
    };

    deleteFeed(id){
        this.http.delete(`/rss/${id}`, { feedId: id }).then((response) => {
            this.feeds = response.data;
        });
    };

    setNewFeedName(id, name) {
        this.http.put(`/rss/${id}`, { feedId: id, name: name }).then((response) => {
            this.feeds = response.data;
        });
    };

    updateAll() {
        this.showWaitbar = true;

        this.http.get('/rss').then((response) => {
            this.feeds = response.data;
            this.showWaitbar = false;
        });
    };

    toggleBookmark(articleId, page, isBookmark, isBookmarkPage, feedId) {
        this.http.put(`/rss/${feedId}/articles/${articleId}`, { isBookmark: isBookmark }).then((response) => {
            if (!response.data) {
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
        let config = {};
        config.params = { page: page };

        this.http.get(`/rss/articles/bookmarks`, config).then((response) => {
            this.articles = response.data.Articles;
            this.articlesCount = response.data.Count;
            this.showArticle = false;
        });
    };

    markReadAll(feedId) {
        let config = {};
        config.params = { isReadAll: true };

        this.http.put(`/rss/${feedId}`, config).then((response) => {
            this.articles = response.data.Articles;
            this.articlesCount = response.data.Count;
        });
    };

    createOpml() {

        return this.http.get('/rss/opml');
    };

    markAsRead(id, feedId, page, isRead, userId) {
        let params = { articleId: id, feedId: feedId, page: page, isRead: isRead, userId: userId };

        this.http.put(`/rss/${feedId}/articles/${id}`, params).then((response) => {
            this.articles = response.data.Articles;
            this.articlesCount = response.data.Count;
        });
    };

    setUnread(id, feedId, page, isRead) {
        this.http.put(`/rss/${feedId}/articles/${id}`, { isUnread: isUnread });
    };
}
RssService.$inject = ['$http'];

angular.module('app').service('rssService', RssService);