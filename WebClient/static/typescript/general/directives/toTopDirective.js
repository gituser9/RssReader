function ToTopDirective() {
    'use strict';

    var layout = '<span><md-button id="go-to-top-btn" class="md-fab" aria-label="To Up" ng-click="goToTop()">'
        + '<i class="material-icons">keyboard_arrow_up</i>'
        + '</md-button></span>';

    return {
        template: layout,
        replace: true,
        restrict: 'E',
        link: function (scope, element) {
            element.on('click', function () {
                $('html,body').scrollTop(0);
            });
        }
    }
}
