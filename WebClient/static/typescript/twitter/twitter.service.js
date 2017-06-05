function TwitterService($http) {
    'use strict';

    var factory = {
        model: {
            News: [],
            Sources: [],
            IsLoad: false,
            IsAll: false
        }
    };

    factory.getPageData = function(id) {
        factory.model.IsLoad = true;
        var config = {};
        config.params = { id: id };

        $http.get("/get-", config).then(function (response) {
            factory.model.News = response.data.News;
            factory.model.Sources = response.data.Groups;
            factory.model.IsLoad = false;
        });
    };

    factory.getNews = function(userId, page) {
        if (factory.model.IsLoad || factory.model.IsAll) {
            return;
        }

        factory.model.IsLoad = true;
        var config = {};
        config.params = { id: userId, page: page };

        $http.get("/get-", config).then(function (response) {
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

    factory.getVkGroups = function(userId) {
        var config = {};
        config.params = { id: userId };

        $http.get("/get-", config).then(function (response) {
            factory.model.Sources = response.data;
        });
    };

    factory.getByFilters = function (filters) {
        var data = {
            GroupId: Number(filters.GroupId)
        };
        $http.post('/get-', data).then(function (response) {
            factory.model.News = response.data;
        });
    };

    return factory;
}
TwitterService.$inject = ['$http'];


