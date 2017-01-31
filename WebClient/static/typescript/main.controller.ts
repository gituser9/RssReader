/// <reference path="_all.ts" />

module main {
    "use strict";

    import IDialogService = angular.material.IDialogService;
    import IToastService = angular.material.IToastService;


    interface IMainScope extends ng.IScope {
        vm: MainController;
        feeds: Feed[];
        articles: Article[];
        article: Article;
        settings: Settings;
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
    }

    class Tab {
        public title: string;
        public article: Article;
    }

    export class MainController {
        public static $inject = [
            "$scope",
            "$timeout",
            "$mdDialog",
            "$mdToast",
            "Upload",
            "mainService"
        ];
        private isBookmark: boolean;
        private isAuth: boolean;
        private currentFeedId: number;
        private userId: number;

        constructor(
            private $scope: IMainScope, 
            private $timeout: ng.ITimeoutService,
            private $mdDialog: IDialogService, 
            private $mdToast: IToastService,
            private $upload: any, 
            private mainService: MainService
        ) {
            $scope.vm = this;
            $scope.searchFeed = 0;
            this.$scope.currentPage = 1;
            this.isBookmark = false;
            this.isAuth = false;
            this.$scope.showArticleCount = true;
            this.currentFeedId = 0;
            this.$scope.tabs = [];

            $scope.$watch(() => {
                this.$scope.feeds = mainService.feeds;
                this.$scope.articles = mainService.articles;
                this.$scope.article = mainService.article;
                this.$scope.showWaitBar = mainService.showWaitbar;
                this.$scope.showArticle = mainService.showArticle;
                this.$scope.articlesCount = mainService.articlesCount;
                this.userId = mainService.currentUserId;
            });            

            let storage = window.localStorage;
            let userStr = storage.getItem("RssReaderUser");

            if (userStr != null) {
                let user = <User> JSON.parse(userStr);

                this.mainService.getAll(user.Id);

                mainService.currentUserId = user.Id;
                this.isAuth = true;
                this.$scope.username = user.Name;

                // this.$timeout(() => { this.mainService.getAll(user.Id) }, 30);  
            } else {
                // modal for auth
                this.openAuthModal();
            }            
        }

        public getAll(): void {
            this.mainService.getAll(this.userId);
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

            this.mainService.getArticles(feed.Feed.Id, this.$scope.currentPage);
        }

        public stepBack(): void {
            this.mainService.showArticle = false;
        }

        public getArticle(article: Article): void {
            this.mainService.getArticle(article.Id);
            this.setRead();
        }

        public updateAll(): void {
            this.mainService.updateAll();
        }

        public getArticlesByPage(page: number): void {
            this.$scope.currentPage = page;
             if (this.isBookmark) {
                 this.mainService.getBookmarks(page);
             } else {
                 this.getArticles(this.$scope.currentFeed);
             }
        }

        public search(): void {
            this.$scope.hideMarkReadAll = true;
            this.mainService.search(this.$scope.searchText, this.$scope.searchInBookmark, this.$scope.searchFeed);
            this.$scope.currentFeedTitle = `Search: ${this.$scope.searchText}`;
        }

        // todo: split this and unset
        public setBookmark(articleId: number): void {
            this.mainService.toggleBookmark(
                articleId,
                this.$scope.currentPage,
                true,
                this.isBookmark,
                this.currentFeedId
            );
        }

        public unsetBookmark(articleId: number): void {
            this.mainService.toggleBookmark(
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

            this.mainService.getBookmarks(1);
        }

        public markReadAll(): void {
            this.mainService.markReadAll(this.currentFeedId);

            this.$scope.currentFeed.ArticlesCount = 0;
            this.$scope.currentFeed.ExistUnread = false;
        }

        public toggleAsRead(id: number, isRead: boolean): void {
            this.mainService.markAsRead(id, this.currentFeedId, this.$scope.currentPage, isRead);

            if (isRead) {
                this.setRead();
            } else {
                ++this.$scope.currentFeed.ArticlesCount;
                this.$scope.currentFeed.ExistUnread = true;
            }
        }

        public createOpml(): void {
            console.log(this.userId);
            this.mainService.createOpml(this.userId).then((): void => {
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

        public logout(): void {
            let storage = window.localStorage;
            storage.removeItem("RssReaderUser");

            this.mainService.feeds = [];
            this.mainService.articles = [];
            this.mainService.article = null;
            this.$scope.currentFeedTitle = "";

            this.openAuthModal();
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

            this.mainService.getArticlePromise(id).then((response: ng.IHttpPromiseCallbackArg<Article>): void => {
                tab.article = response.data;

                this.mainService.articles.forEach((item: Article) => {
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
/*
Modals
================================================================================
*/
        public openDelete(rss: Feeds): void {
            let modalData = new ModalData();
            modalData.Feed = rss;
            this.openModal("static/html/modals/deleteModal.html", modalData);
        }

        public openAdd(): void {
            this.openModal("static/html/modals/addModal.html", null);
        }

        public openEditName(rss: Feeds): void {
            let modalData = new ModalData();
            modalData.Feed = rss;
            this.openModal("static/html/modals/editModal.html", modalData);
        }

        public openSettings(): void {
            let storage = window.localStorage;
            let userStr = storage.getItem("RssReaderUser");
            let user = <User> JSON.parse(userStr);

            this.mainService.getSettings(user.Id).then((response: Settings): void => {
                let modalData = new ModalData();
                modalData.Settings = response;
                this.openModal("static/html/modals/settingModal.html", modalData);
            });
        }

        public openAuthModal(): void {
            this.$mdDialog.show({
                controller: ModalController,
                templateUrl: "static/html/modals/authModal.html",
                parent: angular.element(document.body),
                clickOutsideToClose: false,
                locals: {
                    modalData: null
                }
            });
        }

/*
Private
================================================================================
*/
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

        private openModal(url: string, modalData?: ModalData): void {
            this.$mdDialog.show({
                controller: ModalController,
                templateUrl: url,
                parent: angular.element(document.body),
                clickOutsideToClose: true,
                locals: {
                    modalData: angular.copy(modalData)
                }
            });
        }
    }

}
