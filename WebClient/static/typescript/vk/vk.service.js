function VkService ($http) {

    var factory = {
        model: {
            VkNews: [],
            VkGroups: [],
            IsLoad: false,
            IsAll: false
        }
    };

    factory.getPageData = function(id) {
        factory.model.IsLoad = true;
        var config = {};
        config.params = { id: id };

        $http.get("/get-vk-page", config).then(function (response) {
            factory.model.VkNews = response.data.News;
            factory.model.VkGroups = response.data.Groups;
            factory.model.IsLoad = false;
        });
    };

    factory.getVkNews = function(userId, page) {
        if (factory.model.IsLoad || factory.model.IsAll) {
            return;
        }

        factory.model.IsLoad = true;
        var config = {};
        config.params = { id: userId, page: page };

        $http.get("/get-vk-news", config).then(function (response) {
            // factory.model.VkNews = response.data;
            if (response.data.length === 0) {
                factory.model.IsAll = true;
                return;
            }

            for (var i = 0; i < response.data.length; ++i) {
                factory.model.VkNews.push(response.data[i]);
            }

            factory.model.IsLoad = false;
        });
    };

    factory.getVkGroups = function(userId) {
        var config = {};
        config.params = { id: userId };

        $http.get("/get-vk-groups", config).then(function (response) {
            factory.model.VkGroups = response.data;
        });
    };

    factory.loadComments = function (news) {
        var url = "https://api.vk.com/method/wall.getComments";
        var cfg = {
            params: {
                post_id: news.PostId,
                count: 100,
                sort: 'asc',
                owner_id: '-' + news.GroupId
            }
        };
        $http.jsonp(url, cfg).then(function (response) {
            console.log(response);
        });
    };

    factory.getByFilters = function (filters) {
        var data = {
            GroupId: Number(filters.GroupId)
        };
        $http.post('/get-vk-news-by-filters', data).then(function (response) {
            factory.model.VkNews = response.data;
        });
    };

    return factory;
}
VkService.$inject = ["$http"];