/// <reference path="_all.ts" />

module main {
    "use strict";

    import IDialogService = angular.material.IDialogService;


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
        currentFeedTitle: string;
        searchText: string;
        currentLink: string;
        searchFeed: number;
        articlesCount: number;
        currentPage: number;
    }

    export class MainController {
        public static $inject = [
            "$scope",
            "$mdDialog",
            "Upload",
            "mainService"
        ];
        private isBookmark: boolean;
        private currentFeedId: number;

        constructor(private $scope: IMainScope, private $mdDialog: IDialogService, private $upload: any, private mainService: MainService) {
            $scope.vm = this;
            $scope.searchFeed = 0;
            this.$scope.currentPage = 1;

            $scope.$watch(() => {
                this.$scope.feeds = mainService.feeds;
                this.$scope.articles = mainService.articles;
                this.$scope.article = mainService.article;
                this.$scope.showWaitBar = mainService.showWaitbar;
                this.$scope.showArticle = mainService.showArticle;
                this.$scope.articlesCount = mainService.articlesCount;
            });

            this.isBookmark = false;
            this.currentFeedId = 0;

            this.mainService.getAll();
        }

        public getAll(): void {
            this.mainService.getAll();
        }

        public getArticles(feed: Feed): void {
            this.$scope.hideMarkReadAll = false;

            if (feed.Rss.ID != this.currentFeedId) {
                this.$scope.currentPage = 1;
            }

            this.isBookmark = false;
            this.$scope.currentFeed = feed;
            this.currentFeedId = feed.Rss.ID;
            this.$scope.currentFeedTitle = feed.Rss.RssName;

            this.mainService.getArticles(feed.Rss.ID, this.$scope.currentPage);
        }

        public stepBack(): void {
            this.mainService.showArticle = false;
        }

        public getArticle(article: Article): void {
            this.mainService.getArticle(article.ID);
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
            this.mainService.createOpml().then((): void => {
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
                data: { file: file }
            }).success((data: Feed[]) => {
                selfScope.showWaitBar = false;
                selfScope.feeds = data;
            });
        }

/*
Modals
================================================================================
*/    
        public openDelete(rss: Rss): void {
            let modalData = new ModalData();
            modalData.Rss = rss;
            this.openModal("static/modals/deleteModal.html", modalData);
        }

        public openAdd(): void {
            this.openModal("static/modals/addModal.html", null);
        }

        public openEditName(rss: Rss): void {
            let modalData = new ModalData();
            modalData.Rss = rss;
            this.openModal("static/modals/editModal.html", modalData);
        }

        public openSettings(): void {
            this.mainService.getSettings().then((response: Settings): void => {
                let modalData = new ModalData();
                modalData.Settings = response;
                this.openModal("static/modals/settingModal.html", modalData);
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
