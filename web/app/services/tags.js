angular.module('myApp.services.tags', [])
.factory('tags', ['$http','$location', function($http,$location) {
	var tags = {};
	tags.get = function(query) {
		return $http.get('/ws/tags/'+query);
	};

	return tags;
}
]);


