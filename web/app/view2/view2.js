'use strict';

angular.module('myApp.view2', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
	$routeProvider.when('/view2', {
		templateUrl: 'view2/view2.html',
		controller: 'View2Ctrl'
	});
}])

.controller('View2Ctrl', ['$scope','events',function($scope, events) {

	$scope.addEvent= function(){
		var event = {"name": $scope.name, "start": $scope.start, "end": $scope.end};
		events.add(event);
	}
}
]);
