'use strict';

// Declare app level module which depends on views, and components
angular.module('myApp', [
  'ngRoute',
  'myApp.view1',
  'myApp.view2',
  'myApp.services.events',
  'myApp.services.tags',
  'ui.bootstrap.typeahead',
  'ui.bootstrap.tpls',
  'ui.bootstrap.transition',
  'ngTagsInput'
]).
	config(['$routeProvider', function($routeProvider) {
	$routeProvider.otherwise({redirectTo: '/view1'});
}]);
