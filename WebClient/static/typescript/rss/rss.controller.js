
function RssController ($scope, $timeout, $mdDialog, $mdToast, $upload, mainService, rssService){
        $scope.vm = $scope;
        $scope.searchFeed = 0;
        $scope.currentPage = 1;
        $scope.isBookmark = false;
        $scope.isAuth = false;
        $scope.showArticleCount = true;
        $scope.currentFeedId = 0;
        $scope.tabs = [];

        $scope.$watch(function() {
            $scope.feeds = rssService.feeds;
            $scope.articles = rssService.articles;
            $scope.article = rssService.article;
            $scope.showWaitBar = rssService.showWaitbar;
            $scope.showArticle = rssService.showArticle;
            $scope.articlesCount = rssService.articlesCount;
            $scope.userId = mainService.currentUserId;
        });

    $scope.getAll = function() {
        rssService.showWaitBar = true;
        rssService.getAll($scope.userId);
    };

    $scope.getArticles = function(feed) {
        $scope.hideMarkReadAll = false;

        if (feed.Feed.Id != $scope.currentFeedId) {
            $scope.currentPage = 1;
        }

        $scope.isBookmark = false;
        $scope.currentFeed = feed;
        $scope.currentFeedId = feed.Feed.Id;
        $scope.currentFeedTitle = feed.Feed.Name;

        rssService.getArticles(feed.Feed.Id, $scope.currentPage, $scope.userId);
    };

    $scope.stepBack = function() {
        rssService.showArticle = false;
    };

    $scope.getArticle = function(article) {
        rssService.getArticle(article.Id);
        $scope.setRead();
    };

    $scope.updateAll = function() {
        rssService.updateAll();
    };

    $scope.getArticlesByPage = function(page) {
        $scope.currentPage = page;
            if ($scope.isBookmark) {
                rssService.getBookmarks(page);
            } else {
                $scope.getArticles($scope.currentFeed);
            }
    };

    $scope.search = function() {
        $scope.hideMarkReadAll = true;
        rssService.search($scope.searchText, $scope.searchInBookmark, $scope.searchFeed);
        $scope.currentFeedTitle = 'Search: ' + $scope.searchText;
    };

    // todo: split $scope and unset
    $scope.setBookmark = function(articleId) {
        rssService.toggleBookmark(
            articleId,
            $scope.currentPage,
            true,
            $scope.isBookmark,
            $scope.currentFeedId
        );
    };

    $scope.unsetBookmark = function(articleId) {
        rssService.toggleBookmark(
            articleId,
            $scope.currentPage,
            false,
            $scope.isBookmark,
            $scope.currentFeedId
        );
    };

    /* = functiontoggleBookmark(article: app.services.Article) {
        mainService.toggleBookmark(
            article.ID,
            $scope.currentPage,
            !article.IsRead,
            $scope.isBookmark,
            $scope.currentFeedId
        );
    }*/

    $scope.getBookmarks = function() {
        $scope.isBookmark = true;
        $scope.hideMarkReadAll = true;
        $scope.currentFeedTitle = "Bookmarks";

        rssService.getBookmarks(1);
    };

    $scope.markReadAll = function() {
        rssService.markReadAll($scope.currentFeedId, $scope.userId);

        $scope.currentFeed.ArticlesCount = 0;
        $scope.currentFeed.ExistUnread = false;
    };

    $scope.toggleAsRead = function(id, isRead) {
        rssService.markAsRead(id, $scope.currentFeedId, $scope.currentPage, isRead, $scope.userId);

        if (isRead) {
            $scope.setRead();
        } else {
            ++$scope.currentFeed.ArticlesCount;
            $scope.currentFeed.ExistUnread = true;
        }
    };

    $scope.createOpml = function() {
        console.log($scope.userId);
        rssService.createOpml($scope.userId).then(function() {
            var a = document.createElement("a");
            a.download = name;
            a.href = 'static/rss.opml';

            document.querySelector("#export-link").appendChild(a);
            a.addEventListener("click", function() {
                a.parentNode.removeChild(a);
            });

            a.click();
        });
    };

    $scope.openUploadFile = function() {
        var elem = document.querySelector('#upload-opml');

        if (elem && document.createEvent) {
            var evt = document.createEvent("MouseEvents");
            evt.initEvent("click", true, false);
            elem.dispatchEvent(evt);
        }
    };

    $scope.uploadOpml = function(file) {
        $scope.showWaitBar = true;
        var selfScope = $scope;

        $scope.$upload.upload({
            url: 'upload-opml',
            data: { file: file, userId: $scope.userId }
        }).success(function(data) {
            selfScope.showWaitBar = false;
            selfScope.feeds = data;
        });
    };

    $scope.addTab = function(id, title) {
        // if tab exists - show toast and return
        for (var i = 0; i < $scope.tabs.length; ++i) {
            if ($scope.tabs[i].article.Id == id) {
                $mdToast.show(
                    $mdToast.simple()
                        .textContent('Tab already exists')
                        .position("top end")
                        .hideDelay(1500)
                );
                return;
            }
        }

        // add new tab
        var tab = {};
        tab.title = title;

        $scope.tabs.push(tab);

        rssService.getArticlePromise(id).then(function(response) {
            tab.article = response.data;


            rssService.articles.forEach(function(item) {
                if (item.Id == id) {
                    if (!item.IsRead) {
                        $scope.setRead();
                    }
                    item.IsRead = true;
                    return;
                }
            });
        });
    };

    $scope.removeTab = function(tab) {
        var index = $scope.tabs.indexOf(tab);
        $scope.tabs.splice(index, 1);
    };

    $scope.removeAllTabs = function() {
        $scope.tabs = [];
    };

    $scope.showPreview = function(article) {
        if (!article.IsRead) {
            $scope.setRead();
        }
        article.IsRead = true;

        rssService.getArticlePromise(article.Id).then(function(response) {
            article.Body = response.data.Body;
            article.Link = response.data.Link;
        });
    };

    $scope.hidePreview = function(article) {
        article.Body = "";
    };

    /* Modals
    =========================================================== */

    $scope.openDelete = function(rss) {
        var modalData = {};
        modalData.Feed = rss;
        mainService.openModal("deleteModal.html", RssModalController, modalData);
    };

    $scope.openAdd = function() {
        mainService.openModal("addModal.html", RssModalController, null);
    };

    $scope.openEditName = function(rss) {
        var modalData = {};
        modalData.Feed = rss;
        mainService.openModal("editModal.html", RssModalController, modalData);
    };

    /* Private
    ================================================================================ */
    $scope.setRead = function() {
        if ($scope.isBookmark) {
            return;
        }

        if ($scope.currentFeed.ArticlesCount > 0) {
            --$scope.currentFeed.ArticlesCount;
        }

        if ($scope.currentFeed.ArticlesCount === 0) {
            $scope.currentFeed.ExistUnread = false;
        }
    }
}
RssController.$inject = [
    "$scope",
    "$timeout",
    "$mdDialog",
    "$mdToast",
    "Upload",
    "mainService",
    "rssService"
];
