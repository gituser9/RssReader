class RssModalController {
    constructor($scope, $mdDialog, rssService, modalData) {
        this.$scope = $scope
        this.$mdDialog = $mdDialog
        this.rssService = rssService
        this.modalData = modalData
        this.$scope.vm = this;
        this.$scope.feedUrl = "";

        if (modalData !== null) {
            this.$scope.modalData = modalData;
        }
    }

    updateFeedName() {
        if (!this.$scope.modalData.Feed.Name || !this.$scope.modalData.Feed.Name.trim().length) {
            return;
        }
        this.rssService.setNewFeedName(this.$scope.modalData.Feed.Id, this.$scope.modalData.Feed.Name);
        this.cancel();
        this.rssService.getAll()
    };

    hide() {
        this.$mdDialog.hide();
    };

    cancel() {
        this.$mdDialog.cancel();
    };

    addFeed() {
        if (!this.$scope.feedUrl || !this.$scope.feedUrl.trim().length) {
            return;
        }
        this.rssService.addFeed(this.$scope.feedUrl);
        this.cancel();
    };

    delete() {
        this.rssService.delete(this.$scope.modalData.Feed.Id);
        this.$scope.cancel();
    };

    toggleUnread() {
        this.rssService.setUnread(this.$scope.modalData.Settings.UnreadOnly);
        this.$scope.cancel();
    };
}
RssModalController.$inject = [
    "$scope",
    "$mdDialog",
    "rssService",
    "modalData"
];