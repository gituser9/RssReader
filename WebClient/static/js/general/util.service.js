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

            fetch('/users/refresh', {
                    method: 'PUT',
                    body: JSON.stringify({ 'token': rtoken })
                })
                .then(response => response.json())
                .then(data => {
                    localStorage.setItem('rtoken', data.refresh_token)
                    localStorage.setItem('token', data.token)

                    let cfg = response.config

                    fetch(cfg.url, {
                            method: cfg.method,
                            headers: {
                                'Authorization': 'Bearer ' + data.token
                            },
                            body: rdata === null ? null : JSON.stringify(rdata)
                        })
                        .then(response => {
                            if (response.ok) {
                                if (callback !== null && callback !== undefined) {
                                    try {
                                        response.json()
                                            .then(jsonData => callback(jsonData))
                                    } catch (e) {
                                        callback(null);
                                    }
                                }
                            } else {
                                this.logout()
                            }
                        })
                })
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