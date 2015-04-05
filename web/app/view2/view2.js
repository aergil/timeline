'use strict';

angular.module('myApp.view2', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
		$routeProvider.when('/view2', {
				templateUrl: 'view2/view2.html',
				controller: 'View2Ctrl'
		});
}])

.controller('View2Ctrl', ['$scope','events','tags',function($scope, events,tags) {
		$scope.getEvents = function(val){return events.getByName(val);};
		$scope.asyncEvent = undefined;
		$scope.tags = [];

		$scope.create = function(){
				$scope.asyncEvent = {name:"new",ponctuels:[]};
		}

		$scope.addEvent= function(){
				tags = [];
				angular.forEach($scope.tags, function(value, key) {
						this.push(value.name);
				}, tags);
				$scope.asyncEvent.tags = tags;
				events.add($scope.asyncEvent);
		}

		$scope.addPonctuel = function(){
				$scope.asyncEvent.ponctuels.push({});
		}

		$scope.deletePonctuel = function(p){
				var index = $scope.asyncEvent.ponctuels.indexOf(p);
				$scope.asyncEvent.ponctuels.splice(index,1);
		}

		$scope.loadTags = function(query){
				return tags.get(query);
		}

		$scope.typeaheadOnSelect = function ($item, $model, $label){
				$scope.tags = [];
				angular.forEach($item.tags,function(value,key){
						this.push({"name":value});
				},$scope.tags)
		}

}
]);
