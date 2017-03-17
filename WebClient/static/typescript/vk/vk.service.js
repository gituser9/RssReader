function VkService ($http) {

    var factory = {
        model: {
            VkNews: [],
            VkGroups: []
        }
    };

    factory.getPageData = function(id) {
        var config = {};
        config.params = { id: id };

        $http.get("/get-vk-page", config).then(function (response) {
            factory.model.VkNews = response.data.News;
            factory.model.VkGroups = response.data.Groups;
        });
    };

    factory.getVkNews = function(userId) {
        var config = {};
        config.params = { id: userId };

        $http.get("/get-vk-news", config).then(function (response) {
            console.log(response);
            factory.model.VkNews = response.data;
        });
    };

    factory.getVkGroups = function(userId) {
        var config = {};
        config.params = { id: userId };

        $http.get("/get-vk-groups", config).then(function (response) {
            factory.model.VkGroups = response.data;
        });
    };

    return factory;
}
VkService.$inject = ["$http"];