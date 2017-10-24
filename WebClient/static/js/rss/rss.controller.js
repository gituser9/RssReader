class RssController {
    constructor($scope, $timeout, $mdToast, Upload, mainService, rssService) {
        this.$scope = $scope;
        this.$timeout = $timeout;
        this.$mdToast = $mdToast;
        this.Upload = Upload;
        this.mainService = mainService;
        this.rssService = rssService;

        this.$scope.currentPage = 1;
        this.$scope.isBookmark = false;
        this.$scope.isAuth = false;
        this.$scope.showArticleCount = true;
        this.$scope.currentFeedId = 0;
        this.$scope.tabs = [];
        this.$scope.filters = {
            searchText: '',
            searchFeed: 0,
            searchInBookmark: false
        };

        this.$scope.$watch(() => {
            this.$scope.feeds = this.rssService.feeds;
            this.$scope.articles = this.rssService.articles;
            this.$scope.article = this.rssService.article;
            this.$scope.showWaitBar = this.rssService.showWaitBar;
            this.$scope.showArticle = this.rssService.showArticle;
            this.$scope.articlesCount = this.rssService.articlesCount;
            this.$scope.userId = this.mainService.currentUserId;
        });
    }

    getAll() {
        this.rssService.showWaitBar = true;
        /*this.$timeout(() => {
            rssService.getAll($scope.userId);
        }, 3000);*/
        this.rssService.getAll(this.$scope.userId);
    };

    getArticles(feed) {
        this.$scope.hideMarkReadAll = false;

        if (feed.Feed.Id !== this.$scope.currentFeedId) {
            this.$scope.currentPage = 1;
        }

        this.$scope.isBookmark = false;
        this.$scope.currentFeed = feed;
        this.$scope.currentFeedId = feed.Feed.Id;
        this.$scope.currentFeedTitle = feed.Feed.Name;

        this.rssService.getArticles(feed.Feed.Id, this.$scope.currentPage, this.$scope.userId);
    };

    stepBack() {
        this.rssService.showArticle = false;
    };

    getArticle(article) {
        this.rssService.getArticle(article.Id);
        this.setRead();
    };

    updateAll() {
        this.rssService.updateAll();
    };

    getArticlesByPage(page) {
        this.$scope.currentPage = page;

        if (this.$scope.isBookmark) {
            this.rssService.getBookmarks(page);
        } else {
            this.$scope.getArticles(this.$scope.currentFeed);
        }
    };

    search() {
        this.$scope.hideMarkReadAll = true;

        this.rssService.search(this.$scope.filters);
        // this.rssService.search(this.$scope.searchText, false, this.$scope.searchFeed);
        this.$scope.currentFeedTitle = 'Search: ' + this.$scope.filters.searchText;
    };

    // todo: split this.$scope and unset
    setBookmark(articleId) {
        this.rssService.toggleBookmark(
            articleId,
            this.$scope.currentPage,
            true,
            this.$scope.isBookmark,
            this.$scope.currentFeedId
        );
    };

    unsetBookmark(articleId) {
        this.rssService.toggleBookmark(
            articleId,
            this.$scope.currentPage,
            false,
            this.$scope.isBookmark,
            this.$scope.currentFeedId
        );
    };

    getBookmarks() {
        this.$scope.isBookmark = true;
        this.$scope.hideMarkReadAll = true;
        this.$scope.currentFeedTitle = "Bookmarks";

        this.rssService.getBookmarks(1);
    };

    markReadAll() {
        this.rssService.markReadAll(this.$scope.currentFeedId, this.$scope.userId);

        this.$scope.currentFeed.ArticlesCount = 0;
        this.$scope.currentFeed.ExistUnread = false;
    };

    markReadAllById(id) {
        this.rssService.markReadAll(id, this.$scope.userId);

        this.$scope.currentFeed.ArticlesCount = 0;
        this.$scope.currentFeed.ExistUnread = false;
    };

    toggleAsRead(id, isRead) {
        this.rssService.markAsRead(id, this.$scope.currentFeedId, this.$scope.currentPage, isRead, this.$scope.userId);

        if (isRead) {
            this.setRead();
        } else {
            ++this.$scope.currentFeed.ArticlesCount;
            this.$scope.currentFeed.ExistUnread = true;
        }
    };

    createOpml() {
        this.rssService.createOpml(this.$scope.userId).then(() => {
            let a = document.createElement("a");
            a.download = name;
            a.href = 'static/rss.opml';

            document.querySelector("#export-link").appendChild(a);
            a.addEventListener("click", function() {
                a.parentNode.removeChild(a);
            });

            a.click();
        });
    };

    openUploadFile() {
        let elem = document.querySelector('#upload-opml');

        if (elem && document.createEvent) {
            let evt = document.createEvent("MouseEvents");
            evt.initEvent("click", true, false);
            elem.dispatchEvent(evt);
        }
    };

    uploadOpml(file) {
        this.rssService.showWaitBar = true;
        this.rssService.articles = [];
        this.$scope.currentFeedTitle = '';

        this.$upload.upload({
            url: 'upload-opml',
            data: { file: file, userId: this.$scope.userId }
        }).success(function(data) {
            this.rssService.showWaitBar = false;
            this.rssService.feeds = data;
        });
    };

    addTab(id, title) {
        // if tab exists - show toast and return
        for (let tab of this.$scope.tabs) {
            if (tab.article.Id === id) {
                this.$mdToast.show(
                    this.$mdToast.simple()
                        .textContent('Tab already exists')
                        .position("top end")
                        .hideDelay(1500)
                );
                return;
            }
        }

        // add new tab
        let tab = {};
        tab.title = title;
        this.$scope.tabs.push(tab);

        this.rssService.getArticlePromise(id).then((response) => {
            tab.article = response.data;

            for (let article of this.rssService.articles) {
                if (article.Id === id) {
                    if (!article.IsRead) {
                        this.setRead();
                    }
                    article.IsRead = true;
                    return;
                }
            }
        });
    };

    removeTab(tab) {
        let index = this.$scope.tabs.indexOf(tab);
        this.$scope.tabs.splice(index, 1);
    };

    removeAllTabs() {
        this.$scope.tabs = [];
    };

    showPreview(article) {
        if (!article.IsRead) {
            this.setRead();
        }
        article.IsRead = true;

        this.rssService.getArticlePromise(article.Id).then((response) => {
            article.Body = response.data.Body;
            article.Link = response.data.Link;
        });
    };

    hidePreview(article) {
        article.Body = "";
    };

    /* Modals
    =========================================================== */

    openDelete(rss) {
        let modalData = {};
        modalData.Feed = rss;
        this.mainService.openModal("deleteModal.html", RssModalController, modalData);
    };

    openAdd() {
        this.mainService.openModal("addModal.html", RssModalController, null);
    };

    openEditName(rss) {
        let modalData = {};
        modalData.Feed = rss;
        this.mainService.openModal("editModal.html", RssModalController, modalData);
    };

    /* Private
    ================================================================================ */
    setRead() {
        if (this.$scope.isBookmark) {
            return;
        }

        if (this.$scope.currentFeed.ArticlesCount > 0) {
            --this.$scope.currentFeed.ArticlesCount;
        }

        if (this.$scope.currentFeed.ArticlesCount === 0) {
            this.$scope.currentFeed.ExistUnread = false;
        }
    }
}
RssController.$inject = [
    "$scope",
    "$timeout",
    "$mdToast",
    "Upload",
    "mainService",
    "rssService"
];

angular.module('app').controller('rssCtrl', RssController);