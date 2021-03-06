'use strict';

angular.module('myApp.view1', ['ngRoute'])

.config(['$routeProvider', function ($routeProvider) {
		$routeProvider.when('/view1', {
				templateUrl: 'view1/view1.html',
				controller: 'View1Ctrl'
		});
}])

.controller('View1Ctrl', ['$scope','events','tags', function ($scope,events,tags) {

		$scope.zoom = 1;
		$scope.dateMilieu = 1000;
		$scope.startDate = 0;
		$scope.endDate = 2000;
		$scope.startDateScreen = 0;
		$scope.endDateScreen = 2000;

		var refStartDate = 0;
		var refEndDate = 2000;
		var panelWidth = $("#timenav-background").width();

		var ecartMax = 2000;
		var startDate = $scope.dateMilieu;
		var endDate = $scope.endDate;
		var interval = endDate - startDate;
		var intervalPx = panelWidth / interval;
		var offsetLeft = $("#content").width() / 2 - 3;

		var $elements = $("#elements");
		var $ponctuels = $("#ponctuels");
		var $navYears = $("#nav-years");
		var $navigation = $("#navigation");

		var topMin = 2;
		var topSize = 24;
		var tops = [];

		var ponctuelTopMin = 170;
		var ponctuelTopSize = 36;
		var ponctuelRelativeTops = [];

		var elementsJson = [];

		$scope.tags = [];
		$scope.loadTags = function(query){
				return tags.get(query);
		};
		$scope.onTagChanged = function(){
				init();	
		};

		function init(){
				var tags = [];
				angular.forEach($scope.tags,function(value,key){
						this.push(value.name);
				},tags)

				var callback =  function(data){elementsJson = data;$scope.recalcul();};
				events.getWithTags(startDate,endDate,tags,callback);
		}


		function traceDates() {
				var moduloInterval = 1;
				var minInter = 60;//px
				if (intervalPx < minInter){
						moduloInterval = Math.floor(minInter / intervalPx);
				}

				for (var i = 0; i < interval; i++) {
						if(i % moduloInterval == 0){
								var position = intervalPx * i + offsetLeft;
								var year = parseInt(startDate) + i;
								$navYears.append("<div style='left:" + position + "px'>" + year + "</div>");
						}
				}
		}

		function calculTop(left, width) {
				var top = topMin;
				var i = 1;
				var right = left + width;

				var topsInRange = tops.filter(function (x) { return (x.left < left && x.right > left ) || (x.left >= left && x.left < right) });
				while (topsInRange.filter(function (x) { return x.top == top }).length > 0) {
						top = topMin + ((topMin + topSize) * i);
						i++;
				}

				tops.push({top: top, left: left, right: right});
				return top;
		}

		function calculPonctuelRelativeTop(left) {
				var relativeTop = 0;
				var i = 1;
				var right = left + 50;

				var relativeTopsInRange = ponctuelRelativeTops.filter(function (x) { return (x.left < left && x.right > left ) || (x.left >= left && x.left < right) });
				while (relativeTopsInRange.filter(function (x) { return x.relativeTop == relativeTop }).length > 0) {
						relativeTop = ponctuelTopSize * i;
						i++;
				}

				ponctuelRelativeTops.push({relativeTop: relativeTop, left: left, right: right});
				return relativeTop;
		}

		function traceElementsJson() {
				if(elementsJson.length == 0) return;

				var color = 0;
				var elementsFiltered = elementsJson.filter(function(x){return x.start > startDate && x.end < endDate })
				$.each(elementsFiltered, function (index, value) {
						var left = offsetLeft + intervalPx * (value.start - startDate);
						var width = intervalPx * (value.end - value.start);
						var top = calculTop(left, width);
						color = (color + 1) % 5 + 1;
						$elements.append('<div class="color{0}" style="width:{1}px;left:{2}px;top:{3}px;">{4}</div>'
										 .format(color, width, left, top, value.name));

										 $.each(value.ponctuels, function (index, value) {
												 var position = offsetLeft + intervalPx * (value.date - startDate);
												 var paddingTop = ponctuelTopMin - top + calculPonctuelRelativeTop(left);
												 $ponctuels.append('<div class="ponctuel" style="left:{0}px;top:{1}px;padding-top:{2}px"><div class="date" >{3}</div>{4}</div>'
																   .format(position, top + topSize - 8, paddingTop, value.date, value.description));
										 });
				});
		}

		function scrollToMiddle() {
				$navigation.animate({scrollLeft: panelWidth / 2}, 500);
				$navigation.mousewheel(function (event, delta) {
						this.scrollLeft -= (delta * 100);
						event.preventDefault();

				});
		}

		$scope.recalcul = function() {
				interval = endDate - startDate;
				intervalPx = panelWidth / interval;

				tops = [];
				ponctuelRelativeTops = [];
				$elements.empty();
				$ponctuels.empty();
				$navYears.empty();

				traceDates();
				traceElementsJson();
				scrollToMiddle();
		}

		$scope.changeDateMilieu = function (){
				var zoom = $scope.zoom;
				var ecartSelonZoom = (ecartMax/2 * zoom);
				var dateMilieu = parseInt($scope.dateMilieu);
				var ecartSelonZoom = parseInt(ecartSelonZoom);

				startDate =  dateMilieu - ecartSelonZoom;
				endDate = dateMilieu + ecartSelonZoom;

				$scope.startDate = startDate;
				$scope.endDate = endDate;

				init();
		}

		$scope.changeBoundaryDate = function () {
				startDate = parseInt($scope.startDate == "" ? 0 : $scope.startDate);
				endDate = parseInt($scope.endDate == "" ? 0 : $scope.endDate);

				$scope.dateMilieu = (endDate + startDate) / 2;
				$scope.zoom = (endDate - startDate) / ecartMax;

				init();
		}

		$scope.changeBoundaryDateScreen = function () {
				var screenWidth = screen.width;
				var startDateScreen = parseInt($scope.startDateScreen == "" ? 0 : $scope.startDateScreen);
				var endDateScreen = parseInt($scope.endDateScreen == "" ? 0 : $scope.endDateScreen);

				var ratio = panelWidth / screenWidth;
				var newEcart = Math.ceil((endDateScreen - startDateScreen) * ratio);
				var newMilieu = Math.ceil(startDateScreen + ((endDateScreen - startDateScreen)/2))

				$scope.startDate = newMilieu - (newEcart/2);
				$scope.endDate = newMilieu + (newEcart/2);

				$scope.changeBoundaryDate();
		}



		$scope.changeBoundaryDateScreen();

}]);

