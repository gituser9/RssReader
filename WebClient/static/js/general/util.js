class UtilService {
    constructor($http, $mdDialog) {
        this.http = $http
        this.mdDialog = $mdDialog
        this.modalUrl = "static/html/modals/"
    }

    refreshToken(response, rdata, callback) {
        if (response.status === 403) {
            let rtoken = localStorage.getItem('rtoken')

            if (rtoken === null || rtoken === '') {
                this.logout()
                return
            }

            localStorage.setItem('rtoken', '')
            let xhr = new XMLHttpRequest();
            xhr.open('PUT', '/users/refresh', false);
            xhr.send(JSON.stringify({ 'token': rtoken }));

            if (xhr.status !== 200) {
                this.logout()
                return
            }

            let data = JSON.parse(xhr.responseText)
            localStorage.setItem('rtoken', data.refresh_token)
            localStorage.setItem('token', data.token)

            let cfg = response.config
            let rxhr = new XMLHttpRequest();
            rxhr.open(cfg.method, cfg.url, false);
            rxhr.onreadystatechange = () => {
                if (rxhr.readyState !== 4) {
                    return;
                }
                if (rxhr.status >= 400) {
                    this.logout()
                } else {
                    if (callback !== null && callback !== undefined) {
                        try {
                            callback(JSON.parse(rxhr.responseText));
                        } catch (e) {
                            callback(null);
                        }
                    }
                }
            }
            rxhr.send(rdata === null ? null : JSON.stringify(rdata))

        }
        if (callback !== null) {
            callback(response.data)
        }
    }

    httpGet(url, callback) {
        this.http.defaults.headers.common.Authorization = 'Bearer ' + localStorage.getItem('token');
        this.http.get(url).then(
            (response) => {
                callback(response.data)
            },
            (response) => {
                this.refreshToken(response, null, callback)
            },
        )
    }

    httpPost(url, data, callback) {
        this.http.defaults.headers.common.Authorization = 'Bearer ' + localStorage.getItem('token');
        this.http.post(url, data).then((response) => {
            this.refreshToken(response, data, callback)
        })
    }

    httpPut(url, data, callback) {
        this.http.defaults.headers.common.Authorization = 'Bearer ' + localStorage.getItem('token');
        this.http.put(url, data).then((response) => {
            this.refreshToken(response, data, callback)
        })
    }

    httpDelete(url, callback) {
        this.http.defaults.headers.common.Authorization = 'Bearer ' + localStorage.getItem('token');
        this.http.delete(url).then((response) => {
            this.refreshToken(response, null, callback)
        })
    }

    openModal(template, ctrl, modalData) {
        return this.mdDialog.show({
            controller: ctrl,
            templateUrl: this.modalUrl + template,
            parent: angular.element(document.body),
            clickOutsideToClose: true,
            locals: {
                modalData: angular.copy(modalData)
            }
        });
    };

    logout() {
        let storage = window.localStorage;
        storage.removeItem("token");
        storage.removeItem("rtoken");

        // emit event?
        this.openModal("authModal.html", ModalController, null);
    };
}

UtilService.$inject = ['$http', '$mdDialog']

angular.module('app').service('utilService', UtilService);