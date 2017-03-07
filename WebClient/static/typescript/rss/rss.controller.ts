import { RssModalController } from './rss.modal.controller';
import IDialogService = angular.material.IDialogService;
import IToastService = angular.material.IToastService;

import { User, ModalData } from '../models/generalModels';
import { Feed, Feeds, Article, ArticleData, Tab } from './rss.models';
import { MainService } from '../main.service';
import { RssService } from './rss.service';


interface IRssScope extends ng.IScope {
        vm: RssController;
        feeds: Feed[];
        articles: Article[];
        article: Article;
        // settings: Settings;
        currentFeed: Feed;
        showWaitBar: boolean;
        showArticle: boolean;
        hideMarkReadAll: boolean;
        searchInBookmark: boolean;
        showArticleCount: boolean;
        currentFeedTitle: string;
        searchText: string;
        currentLink: string;
        username: string;
        searchFeed: number;
        articlesCount: number;
        currentPage: number;
        currentUserId: number;
        tabs: Tab[];
        // sources: Sources;
        // currentSource: Sources;
    }


export class RssController {
    public static $inject = [
        "$scope",
        "$timeout",
        "$mdDialog",
        "$mdToast",
        "Upload",
        "mainService",
        "rssService"
    ];

    private isBookmark: boolean;
    private isAuth: boolean;
    private currentFeedId: number;
    private userId: number;

    constructor(
        private $scope: IRssScope, 
        private $timeout: ng.ITimeoutService,
        private $mdDialog: IDialogService, 
        private $mdToast: IToastService,
        private $upload: any, 
        private mainService: MainService,
        private rssService: RssService
    ){
        $scope.vm = this;
        $scope.searchFeed = 0;
        this.$scope.currentPage = 1;
        this.isBookmark = false;
        this.isAuth = false;
        this.$scope.showArticleCount = true;
        this.currentFeedId = 0;
        this.$scope.tabs = [];

        $scope.$watch(() => {
            this.$scope.feeds = rssService.feeds;
            this.$scope.articles = rssService.articles;
            this.$scope.article = rssService.article;
            this.$scope.showWaitBar = rssService.showWaitbar;
            this.$scope.showArticle = rssService.showArticle;
            this.$scope.articlesCount = rssService.articlesCount;
            this.userId = mainService.currentUserId;
        });            

        let storage = window.localStorage;
        let userStr = storage.getItem("RssReaderUser");

        if (userStr != null) {
            let user = <User> JSON.parse(userStr);

            // this.mainService.settings = user.Settings;
            this.mainService.updateSettings(user.Id);
            this.rssService.getAll(user.Id);

            mainService.currentUserId = user.Id;

            this.isAuth = true;
            this.$scope.username = user.Name;

            // this.$timeout(() => { this.mainService.getAll(user.Id) }, 30);  
        } else {
            // modal for auth
            this.mainService.openAuthModal();
        }
    }

    public getAll(): void {
        this.rssService.getAll(this.userId);
    }

    public getArticles(feed: Feed): void {
        this.$scope.hideMarkReadAll = false;

        if (feed.Feed.Id != this.currentFeedId) {
            this.$scope.currentPage = 1;
        }

        this.isBookmark = false;
        this.$scope.currentFeed = feed;
        this.currentFeedId = feed.Feed.Id;
        this.$scope.currentFeedTitle = feed.Feed.Name;

        this.rssService.getArticles(feed.Feed.Id, this.$scope.currentPage);
    }

    public stepBack(): void {
        this.rssService.showArticle = false;
    }

    public getArticle(article: Article): void {
        this.rssService.getArticle(article.Id);
        this.setRead();
    }

    public updateAll(): void {
        this.rssService.updateAll();
    }

    public getArticlesByPage(page: number): void {
        this.$scope.currentPage = page;
            if (this.isBookmark) {
                this.rssService.getBookmarks(page);
            } else {
                this.getArticles(this.$scope.currentFeed);
            }
    }

    public search(): void {
        this.$scope.hideMarkReadAll = true;
        this.rssService.search(this.$scope.searchText, this.$scope.searchInBookmark, this.$scope.searchFeed);
        this.$scope.currentFeedTitle = `Search: ${this.$scope.searchText}`;
    }

    // todo: split this and unset
    public setBookmark(articleId: number): void {
        this.rssService.toggleBookmark(
            articleId,
            this.$scope.currentPage,
            true,
            this.isBookmark,
            this.currentFeedId
        );
    }

    public unsetBookmark(articleId: number): void {
        this.rssService.toggleBookmark(
            articleId,
            this.$scope.currentPage,
            false,
            this.isBookmark,
            this.currentFeedId
        );
    }

    /*public toggleBookmark(article: app.services.Article): void {
        this.mainService.toggleBookmark(
            article.ID,
            this.$scope.currentPage,
            !article.IsRead,
            this.isBookmark,
            this.currentFeedId
        );
    }*/

    public getBookmarks(): void {
        this.isBookmark = true;
        this.$scope.hideMarkReadAll = true;
        this.$scope.currentFeedTitle = "Bookmarks";

        this.rssService.getBookmarks(1);
    }

    public markReadAll(): void {
        this.rssService.markReadAll(this.currentFeedId);

        this.$scope.currentFeed.ArticlesCount = 0;
        this.$scope.currentFeed.ExistUnread = false;
    }

    public toggleAsRead(id: number, isRead: boolean): void {
        this.rssService.markAsRead(id, this.currentFeedId, this.$scope.currentPage, isRead);

        if (isRead) {
            this.setRead();
        } else {
            ++this.$scope.currentFeed.ArticlesCount;
            this.$scope.currentFeed.ExistUnread = true;
        }
    }

    public createOpml(): void {
        console.log(this.userId);
        this.rssService.createOpml(this.userId).then((): void => {
            let a = document.createElement("a");
            a.download = name;
            a.href = 'static/rss.opml';

            document.querySelector("#export-link").appendChild(a);
            a.addEventListener("click", () => {
                a.parentNode.removeChild(a);
            });

            a.click();
        });
    }

    public openUploadFile(): void {
        let elem = document.querySelector('#upload-opml');

        if (elem && document.createEvent) {
            let evt = document.createEvent("MouseEvents");
            evt.initEvent("click", true, false);
            elem.dispatchEvent(evt);
        }
    }

    public uploadOpml(file: any): void {
        this.$scope.showWaitBar = true;
        let selfScope = this.$scope;

        this.$upload.upload({
            url: 'upload-opml',
            data: { file: file, userId: this.userId }
        }).success((data: Feed[]) => {
            selfScope.showWaitBar = false;
            selfScope.feeds = data;
        });
    }

    public addTab(id: number, title: string): void {
        // if tab exists - show toast and return
        for (let i = 0; i < this.$scope.tabs.length; ++i) {
            if (this.$scope.tabs[i].article.Id == id) {
                this.$mdToast.show(
                    this.$mdToast.simple()
                        .textContent('Tab already exists')
                        .position("top end")
                        .hideDelay(3000)
                );
                return;
            }
        }

        // add new tab
        let tab = new Tab();
        tab.title = title;
        
        this.$scope.tabs.push(tab);

        this.rssService.getArticlePromise(id).then((response: ng.IHttpPromiseCallbackArg<Article>): void => {
            tab.article = response.data;

            this.rssService.articles.forEach((item: Article) => {
                if (item.Id == id) {
                    item.IsRead = true;
                    return;
                }
            });
        });
    }

    public removeTab(tab: Tab): void {
        let index = this.$scope.tabs.indexOf(tab);
        this.$scope.tabs.splice(index, 1);
    }

    public removeAllTabs(): void {
        this.$scope.tabs = [];
    }

    public showPreview(article: Article): void {
        article.IsRead = true;

        this.rssService.getArticlePromise(article.Id).then((response: ng.IHttpPromiseCallbackArg<Article>): void => {
            article.Body = response.data.Body;
            article.Link = response.data.Link;
        });
    }

    public hidePreview(article: Article): void {
        article.Body = "";
    }

    /* Modals
    =========================================================== */

    public openDelete(rss: Feeds): void {
        let modalData = new ModalData();
        modalData.Feed = rss;
        this.mainService.openModal("static/html/modals/deleteModal.html", RssModalController, modalData);
    }

    public openAdd(): void {
        this.mainService.openModal("static/html/modals/addModal.html", RssModalController, null);
    }

    public openEditName(rss: Feeds): void {
        let modalData = new ModalData();
        modalData.Feed = rss;
        this.mainService.openModal("static/html/modals/editModal.html", RssModalController, modalData);
    }

    /* Private
    ================================================================================ */
    private setRead(): void {
        if (this.isBookmark) {
            return;
        }

        if (this.$scope.currentFeed.ArticlesCount > 0) {
            --this.$scope.currentFeed.ArticlesCount;
        }

        if (this.$scope.currentFeed.ArticlesCount == 0) {
            this.$scope.currentFeed.ExistUnread = false;
        }
    }
}