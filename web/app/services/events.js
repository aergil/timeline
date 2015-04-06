angular.module('myApp.services.events', [])
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
	
	events.getWithTags = function(begin,end,tags,successCallback) {
		$http.get('/ws/events/'+begin+'/'+end+'/tags/'+tags.join()).
			success(function(data, status, headers, config) {
			successCallback(data);
		}).
			error(function(data, status, headers, config) {
			alert(status);
		});
	};

	events.getByName = function(val){
		return	$http.get('/ws/events/byname/'+val)
		.then(function(response){return response.data;});	
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


