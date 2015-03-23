angular.module('myApp.services', [])
.factory('events', ['$http','$location', function($http,$location) {
	var events = {};
	events.get = function(begin,end,successCallback) {
		$http.get('/ws/events/'+begin+'/'+end).
			success(function(data, status, headers, config) {
			successCallback(data);
		}).
			error(function(data, status, headers, config) {
			alert(status);
		});
	};

	events.add = function(e){
		$http.post('/ws/events', e).
			success(function(data,status,headers,config){
			$location.path('/view/app');
		}).
			error(function(data,status,headers,config){
			alert(status);
		});
	};

	return events;
}
]);


