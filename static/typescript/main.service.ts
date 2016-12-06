/// <reference path="_all.ts" />

module main {
    "use strict";

    export class Feed {
        ArticlesCount: number;
        ExistUnread: boolean;
        Feed: Feeds;
    }

    export class Feeds {
        Id: number;
        Url: string;
        Name: string;
        Articles: Article[];
    }

    export class Article {
        Id: number;
        Title: string;
        Body: string;
        Link: string;
        IsBookMark: boolean;
        IsRead: boolean;
    }

    export class Settings {
        UnreadOnly: boolean;
        MarkSameRead: boolean;
        UpdateMinutes: number;
    }

    export class ArticleData {
        public Articles: Article[];
        public Count: number;
    }

    export class User {
        public Id: number;
        public Name: string;
        public Password: string;
    }

    export class RegistrationData {
        public User: User;
        public Message: string;
    }

    export class MainService {
        public articlesCount: number;
        public currentUserId: number;
        public showArticle: boolean;
        public showWaitbar: boolean;
        public settings: Settings;
        public feeds: Feed[];
        public articles: Article[];
        public article: Article;

        public static $inject = ["$http", "Upload"];

        constructor(private $http: ng.IHttpService) {
            this.$http = $http;
        }

        public getArticles(feedId: number, page: number): void {
            let config: ng.IRequestShortcutConfig = {};
            config.params = { "id": feedId, "page": page };
            
            this.$http.get("/get-articles", config).then((response: ng.IHttpPromiseCallbackArg<ArticleData>): void => {
                this.articles = <Article[]>response.data.Articles;
                this.articlesCount = <number> response.data.Count;
                this.showArticle = false;
            });
        } 

        public search(searchText: string, isBookmark: boolean, feedId: number): void {
            let config: ng.IRequestShortcutConfig = {};
            config.params = { searchString: searchText, isBookmark: isBookmark, feedId: feedId };

            this.$http.get("/search", config).then((response: ng.IHttpPromiseCallbackArg<ArticleData>): void => {
                this.articles = <Article[]>response.data.Articles;
                this.articlesCount = <number> response.data.Count;
                this.showArticle = false;
            });
        }

        public getArticle(id: number): void {
            let config: ng.IRequestShortcutConfig = {};
            config.params = { id: id };

            this.$http.get("/get-article", config).then((response: ng.IHttpPromiseCallbackArg<Article>): void => {
                this.article = response.data;
                this.showArticle = true;

                this.articles.forEach((item: Article) => {
                    if (item.Id == this.article.Id) {
                        item.IsRead = true;
                    }
                });
            });
        }

        public getAll(id: number): void {
            let config: ng.IRequestShortcutConfig = {};
            config.params = { id: id };

            this.$http.get("/get-all", config).then((response: ng.IHttpPromiseCallbackArg<Feed[]>): void => {
                this.feeds = response.data;
            });
        }

        public addFeed(url: string): void {
            this.$http.post("/add-article", { url: url, userId: this.currentUserId }).then((response: ng.IHttpPromiseCallbackArg<Feed[]>): void => {
                this.feeds = response.data;
            });
        }

        public delete(id: number): void {
            this.$http.post('/delete', { feedId: id }).then((response: ng.IHttpPromiseCallbackArg<Feed[]>): void => {
                this.feeds = response.data;
            });
        }

        public setNewFeedName(id: number, name: string): void {
            this.$http.post('/set-new-name', { feedId: id, name: name }).then((response: ng.IHttpPromiseCallbackArg<Feed[]>): void => {
                this.feeds = response.data;
            });
        }

        public updateAll(): void {
            this.showWaitbar = true;

            this.$http.get('/update-all').then((response: ng.IHttpPromiseCallbackArg<Feed[]>): void => {
                this.feeds = response.data;
                this.showWaitbar = false;
            });
        }

        public setSettings(json: string): void {
            this.$http.post('/set-settings', { settings: json });
        }

        public toggleBookmark(articleId: number, page: number, isBookmark: boolean, isBookmarkPage: boolean, feedId: number): void {
            this.$http.post("/toggle-bookmark", { articleId: articleId, page: page, isBookmark: isBookmark }).then((response: ng.IHttpPromiseCallbackArg<boolean>): void => {
                if (!response.data) {
                    return;
                }

                if (isBookmarkPage) {
                    this.getBookmarks(page);
                } else {
                    this.getArticles(feedId, page);
                }
            });
        }

        public getBookmarks(page: number): void {
            let config: ng.IRequestShortcutConfig = {};
            config.params = { page: page };

            this.$http.get("/get-bookmarks", config).then((response: ng.IHttpPromiseCallbackArg<ArticleData>): void => {
                this.articles = <Article[]>response.data.Articles;
                this.articlesCount = <number> response.data.Count;
                this.showArticle = false;
            });
        }
        
        public markReadAll(feedId: number) {
            let config: ng.IRequestShortcutConfig = {};
            config.params = { id: feedId };

            this.$http.get("/mark-read-all", config).then((response: ng.IHttpPromiseCallbackArg<ArticleData>): void => {
                this.articles = <Article[]>response.data.Articles;
                this.articlesCount = <number> response.data.Count;
            });
        }

        // todo: type
        public createOpml(): ng.IHttpPromise<any> {
            return this.$http.get('/create-opml');
        }

        //public setUnread

        public getSettings(): ng.IPromise<Settings> {
            return this.$http.get("/get-settings").then((response: ng.IHttpPromiseCallbackArg<Settings>): Settings => {
                return <Settings> response.data;
            });
        }

        public markAsRead(id: number, feedId: number, page: number, isRead: boolean): void {
            this.$http.post("/toggle-as-read", { articleId: id, feedId: feedId, page: page, isRead: isRead }).then((response: ng.IHttpPromiseCallbackArg<ArticleData>): void => {
                this.articles = response.data.Articles;
                this.articlesCount = response.data.Count;
            });
        }

        public setUnread(isUnread: boolean): void {
            this.$http.post('/toggle-unread', { isUnread: isUnread });
        }

        public auth(username: string, password: string): ng.IHttpPromise<User> {
            return this.$http.post('/auth', { username: username, password: password });
        }

        public registration(username: string, password: string): ng.IHttpPromise<RegistrationData> {
            return this.$http.post('/registration', { username: username, password: password });
        }
    }
}
