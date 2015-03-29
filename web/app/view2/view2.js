'use strict';

angular.module('myApp.view2', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
	$routeProvider.when('/view2', {
		templateUrl: 'view2/view2.html',
		controller: 'View2Ctrl'
	});
}])

.controller('View2Ctrl', ['$scope','events',function($scope, events) {

	$scope.getEvents = function(val){return events.getByName(val);};
	$scope.asyncEvent = undefined;

	$scope.addEvent= function(){
		events.add($scope.asyncEvent);
	}

	$scope.addPonctuel = function(){
		$scope.asyncEvent.ponctuels.push({});
	}
}
]);
