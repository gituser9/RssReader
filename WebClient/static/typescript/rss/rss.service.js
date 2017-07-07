function RssService ($http){

    var factory = {
        articles: [],
        articlesCount: 0,
        showWaitBar: false
    };

    factory.getArticles = function(feedId, page, userId) {
        var config = {};
        config.params = { "id": feedId, "page": page, "userId": userId };

        $http.get("/get-articles", config).then(function(response) {
            factory.articles = response.data.Articles;
            factory.articlesCount = response.data.Count;
            factory.showArticle = false;
        });
    };

    factory.search = function(searchText, isBookmark, feedId) {
        var config = {};
        config.params = { searchString: searchText, isBookmark: isBookmark, feedId: feedId };

        $http.get("/search", config).then(function(response) {
            factory.articles = response.data.Articles;
            factory.articlesCount = response.data.Count;
            factory.showArticle = false;
        });
    };

    factory.getArticle = function(id) {
        var config = {};
        config.params = { id: id };

        $http.get("/get-article", config).then(function(response) {
            factory.article = response.data;
            factory.showArticle = true;

            factory.articles.forEach(function(item) {
                if (item.Id == factory.article.Id) {
                    item.IsRead = true;
                }
            });
        });
    };

    factory.getArticlePromise = function(id) {
        var config = {};
        config.params = { id: id };

        return $http.get("/get-article", config);
    };

    factory.getAll = function(id) {
        var config = {};
        config.params = { id: id };

        $http.get("/get-all", config).then(function(response) {
            factory.feeds = response.data;
            factory.showWaitBar = false;
        });
    };

    factory.addFeed = function(url, userId) {
        $http.post("/add-article", { url: url, userId: userId }).then(function(response) {
            factory.feeds = response.data;
        });
    };

    factory.delete = function(id){
        $http.post('/delete', { feedId: id }).then(function(response) {
            factory.feeds = response.data;
        });
    };

    factory.setNewFeedName = function(id, name) {
        $http.post('/set-new-name', { feedId: id, name: name }).then(function(response) {
            factory.feeds = response.data;
        });
    };

    factory.updateAll = function() {
        factory.showWaitbar = true;

        $http.get('/update-all').then(function(response) {
            factory.feeds = response.data;
            factory.showWaitbar = false;
        });
    };

    factory.toggleBookmark = function(articleId, page, isBookmark, isBookmarkPage, feedId) {
        $http.post("/toggle-bookmark", { articleId: articleId, page: page, isBookmark: isBookmark }).then(function(response) {
            if (!response.data) {
                return;
            }

            if (isBookmarkPage) {
                factory.getBookmarks(page);
            } else {
                factory.getArticles(feedId, page);
            }
        });
    };

    factory.getBookmarks = function(page) {
        var config = {};
        config.params = { page: page };

        $http.get("/get-bookmarks", config).then(function(response) {
            factory.articles = response.data.Articles;
            factory.articlesCount = response.data.Count;
            factory.showArticle = false;
        });
    };

    factory.markReadAll = function(feedId, userId) {
        var config = {};
        config.params = { id: feedId, userId: userId };

        $http.get("/mark-read-all", config).then(function(response) {
            factory.articles = response.data.Articles;
            factory.articlesCount = response.data.Count;
        });
    };

    // todo: type
    factory.createOpml = function(userId) {
        var config = {};
        config.params = { id: userId };

        return $http.get('/create-opml', config);
    };

    factory.markAsRead = function(id, feedId, page, isRead, userId) {
        var params = { articleId: id, feedId: feedId, page: page, isRead: isRead, userId: userId };

        $http.post("/toggle-as-read", params).then(function(response) {
            factory.articles = response.data.Articles;
            factory.articlesCount = response.data.Count;
        });
    };

    factory.setUnread = function(isUnread) {
        $http.post('/toggle-unread', { isUnread: isUnread });
    };

    return factory;
}
RssService.$inject = ["$http"];
