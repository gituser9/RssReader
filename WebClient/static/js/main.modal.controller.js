class ModalController {
    constructor($scope, $mdDialog, mainService, modalData) {
        this.$scope = $scope;
        this.$mdDialog = $mdDialog;
        this.mainService = mainService;
        this.modalData = modalData;
        this.$scope.feedUrl = "";
        this.$scope.vm = this;

        if (modalData !== null) {
            this.$scope.modalData = modalData;
        }
    }

    hide() {
        this.$mdDialog.hide();
    };

    cancel() {
        this.$mdDialog.cancel();
    };

    auth() {
        this.mainService.auth(this.$scope.username, this.$scope.password).then((response) => {
            if (response.data !== null) {
                this.cancel();
                this.mainService.currentUserId = response.data.Id;
                this.mainService.settings = response.data.Settings;
                // this.mainService.getAll(this.mainService.currentUserId);
                this.mainService.updateSettings(response.data.Id);

                let storage = window.localStorage;
                storage.setItem("RssReaderUser", JSON.stringify(response.data));

            }
        });
    };

    registration() {
        this.mainService.registration(this.$scope.username, this.$scope.password).then((response) => {
            if (response.data !== null) {
                let data = response.data;

                if (data.User === null) {
                    this.$scope.errorMessage = data.Message;
                    return
                }

                this.$scope.cancel();

                this.$scope.errorMessage = "";
                this.mainService.currentUserId = data.User.Id;
                this.mainService.settings = data.User.Settings;
                
                let storage = window.localStorage;
                storage.setItem("RssReaderUser", JSON.stringify(data.User));

                this.mainService.openModal("settingModal.html", ModalController, {});
            }
        });
    };


    saveSettings() {
        this.mainService.setSettings(this.$scope.modalData.Settings);
        this.mainService.settings = this.$scope.modalData.Settings;
        
        let storage = window.localStorage;
        let userStr = storage.getItem("RssReaderUser");
        let user = JSON.parse(userStr);
        user.Settings = this.$scope.modalData.Settings;

        storage.setItem("RssReaderUser", JSON.stringify(user));
        this.cancel();
    }
}
ModalController.$inject = [
    "$scope",
    "$mdDialog",
    "mainService",
    "modalData"
];

// angular.module('app').controller('twitterCtrl', ModalController);