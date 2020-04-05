class VkService {
    constructor($http) {
        this.model = {
            VkNews: [],
            VkGroups: [],
            IsLoad: false,
            IsAll: false,
            VkGroupMap: {},
            IsSearch: false
        };
        this.http = $http;
    }

    getPageData() {
        this.model.IsLoad = true;

        this.http.get("/vk/page").then((response) => {
            this.model.VkNews = response.data.News;
            this.model.VkGroups = response.data.Groups;
            this.model.IsLoad = false;

            response.data.Groups.forEach((item) => {
                this.model.VkGroupMap[item.Gid] = {
                    image: item.Image,
                    name: item.Name,
                    link: item.LinkedName
                };
            });
        });
    }

    getVkNews(page) {
        if (this.model.IsLoad || this.model.IsAll || this.model.IsSearch) {
            return;
        }

        this.model.IsLoad = true;
        let config = {};
        config.params = { page: page };

        this.http.get("/vk/news", config).then((response) => {
            this.model.IsLoad = false;

            if (response.data.length === 0) {
                this.model.IsAll = true;
                return;
            }

            for (let item of response.data) {
                this.model.VkNews.push(item);
            }
        });
    };

    getVkGroups(userId) {
        let config = {};
        config.params = { id: userId };

        this.http.get("/vk/groups", config).then((response) => {
            factory.model.VkGroups = response.data;
        });
    };

    loadComments(news) {
        let url = "https://api.vk.com/method/wall.getComments";
        let cfg = {
            params: {
                post_id: news.PostId,
                count: 100,
                sort: 'asc',
                owner_id: '-' + news.GroupId
            }
        };
        this.http.jsonp(url, cfg).then((response) => {
            console.log(response);
        });
    };

    getByFilters(filters) {
        this.model.IsSearch = false;
        let data = {
            GroupId: Number(filters.GroupId)
        };
        this.http.get('/vk/news', data).then((response) => {
            this.model.VkNews = response.data;
        });
    };

    search(searchString, groupId) {
        this.model.IsSearch = true;
        let data = {
            SearchString: searchString,
            GroupId: groupId
        };
        this.http.get('/vk/search', data).then((response) => {
            this.model.VkNews = response.data;
        });
    };
}
VkService.$inject = ['$http'];


angular.module('app').service('vkService', VkService);