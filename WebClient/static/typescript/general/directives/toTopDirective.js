const layout = `<span>
                    <md-button id="go-to-top-btn" class="md-fab" aria-label="To Up" ng-click="goToTop()">
                        <i class="material-icons">keyboard_arrow_up</i>
                    </md-button>
                </span>`;

class ToTopDirective {
    constructor() {
        this.restrict = 'E';
        this.template = layout;
        this.replace = true;
    }

    link(scope, element) {
        element.on('click', function () {
            $('html,body').scrollTop(0);
        });
    }
}

angular.module('app').directive('toTop', () => { return new ToTopDirective() });
