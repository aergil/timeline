angular.module('myApp.services', [])
.factory('events', ['$http', function($http) {
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

     return events;
  }
]);
