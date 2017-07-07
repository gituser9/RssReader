function TwitterService($http) {
    'use strict';

    var factory = {
        model: {
            News: [],
            Sources: [],
            IsLoad: false,
            IsAll: false,
            SourceMap: {}
        }
    };

    factory.getPageData = function(id) {
        factory.model.IsLoad = true;
        var config = {};
        config.params = { id: id };

        $http.get("/get-twitter-page", config).then(function (response) {
            factory.model.News = response.data.News;
            factory.model.Sources = response.data.Sources;
            factory.model.IsLoad = false;

            factory.model.Sources.forEach(function (item) {
                factory.model.SourceMap[item.Id] = {
                    image: item.Image,
                    name: item.Name,
                    link: item.Url,
                    screenName: item.ScreenName
                }
            });

            console.log(factory.model.SourceMap);

        });
    };

    factory.getNews = function(userId, page) {
        if (factory.model.IsLoad || factory.model.IsAll) {
            return;
        }

        factory.model.IsLoad = true;
        var config = {};
        config.params = { id: userId, page: page };

        $http.get("/get-twitter-news", config).then(function (response) {
            if (response.data.length === 0) {
                factory.model.IsAll = true;
                return;
            }

            for (var i = 0; i < response.data.length; ++i) {
                factory.model.News.push(response.data[i]);
            }

            factory.model.IsLoad = false;
        });
    };

    factory.getSources = function(userId) {
        var config = {};
        config.params = { id: userId };

        $http.get("/get-", config).then(function (response) {
            factory.model.Sources = response.data;
        });
    };

    factory.getByFilters = function (filters) {
        factory.model.News = [];
        var data = {
            SourceId: Number(filters.SourceId)
        };
        $http.post('/get-twitter-news-by-filters', data).then(function (response) {

            factory.model.News = response.data;
        });
    };

    factory.search = function (searchString, sourceId) {
        factory.model.IsSearch = true;
        var data = {
            SearchString: searchString,
            SourceId: sourceId
        };
        $http.post('/search-twitter-news', data).then(function (response) {
            factory.model.News = response.data;
        });
    };

    return factory;
}
TwitterService.$inject = ['$http'];


