angular.module('myApp.services', [])
.factory('tags', ['$http','$location', function($http,$location) {
	var tags = {};
	tags.get = function(query,successCallback) {
		return $http.get('/ws/tags/'+query);
	};

	return tags;
}
]);


