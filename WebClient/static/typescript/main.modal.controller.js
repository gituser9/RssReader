function ModalController ($scope, $mdDialog, mainService, modalData) {

    $scope.vm = this;
    $scope.feedUrl = "";

    if (modalData != null) {
        $scope.modalData = modalData;
    }


    $scope.hide = function() {
        $mdDialog.hide();
    };

    $scope.cancel = function() {
        $mdDialog.cancel();
    };

    

    $scope.auth = function() {
        mainService.auth($scope.username, $scope.password).then(function (response) {
            if (response.data != null) {
                $scope.cancel();
                mainService.currentUserId = response.data.Id;
                mainService.settings = response.data.Settings;
                // mainService.getAll(mainService.currentUserId);
                mainService.updateSettings(response.data.Id);

                var storage = window.localStorage;
                storage.setItem("RssReaderUser", JSON.stringify(response.data));
            }
        });
    };

    $scope.registration = function() {
        mainService.registration($scope.username, $scope.password).then(function (response) {
            if (response.data != null) {
                var data = response.data;

                if (data.User == null) {
                    $scope.errorMessage = data.Message;
                    return
                }

                $scope.cancel();

                $scope.errorMessage = "";
                mainService.currentUserId = data.User.Id;
                mainService.settings = data.User.Settings;
                
                var storage = window.localStorage;
                storage.setItem("RssReaderUser", JSON.stringify(data.User));
            }
        });
    };

    $scope.saveSettings = function() {
        mainService.setSettings($scope.modalData.Settings);
        mainService.settings = $scope.modalData.Settings;
        
        var storage = window.localStorage;
        var userStr = storage.getItem("RssReaderUser");
        var user = JSON.parse(userStr);
        user.Settings = $scope.modalData.Settings;

        storage.setItem("RssReaderUser", JSON.stringify(user));
        $scope.cancel();
    }
}
ModalController.$inject = [
    "$scope",
    "$mdDialog",
    "mainService",
    "modalData"
];