import { VkGroup, VkNews } from './vk.models';


export class VkService {
    
    public vkNews: VkNews[];
    public vkGroups: VkGroup[];

    public static $inject = ["$http", "Upload"];

    constructor(private $http: ng.IHttpService) {
        this.$http = $http;
    }

    public getVkNews(userId: number): void {
        let config: ng.IRequestShortcutConfig = {};
        config.params = { id: userId };

        this.$http.get("/get-vk-news", config).then((response: ng.IHttpPromiseCallbackArg<VkNews[]>): void => {
            this.vkNews = <VkNews[]> response.data;
        });
    }

    public getVkGroups(userId: Number): void {
        let config: ng.IRequestShortcutConfig = {};
        config.params = { id: userId };

        this.$http.get("/get-vk-groups", config).then((response: ng.IHttpPromiseCallbackArg<VkGroup[]>): void => {
            this.vkGroups = <VkGroup[]> response.data;
        });
    }
}