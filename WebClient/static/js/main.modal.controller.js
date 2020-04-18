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
            if (response.status === 200) {
                this.cancel();
                let data = response.data;
                let storage = window.localStorage;
                storage.setItem('rtoken', data.refresh_token)
                storage.setItem('token', data.token)

                // this.mainService.getAll();
                this.mainService.updateSettings();
            }
        });
    };

    registration() {
        this.mainService.registration(this.$scope.username, this.$scope.password).then((response) => {
            if (response.status === 200) {
                let data = response.data;

                let storage = window.localStorage;
                storage.setItem('rtoken', data.refresh_token)
                storage.setItem('token', data.token)

                this.cancel();
                this.$scope.errorMessage = "";

                this.mainService.openModal("settingModal.html", ModalController, {});
            }
        });
    };


    saveSettings() {
        this.mainService.setSettings(this.$scope.modalData.Settings);
        this.mainService.settings = this.$scope.modalData.Settings;
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