var app = angular.module('login', ['ngMaterial']);
app.controller('mainCtrl', function($scope,$http,$window) {
	 
	$scope.user={email: '', password: '',remember:true}
	$scope.message={error:''}
    $scope.Send = function() {
    	if($scope.user.email!=''){
    		if($scope.user.password!=''){
		    	$http({
				    url: '/login', 
				    method: "GET",
				    params: {email: $scope.user.email,
				    		 password: $scope.user.password,
				    		 remember: $scope.user.remember	}
				 }).then(function successCallback(response) {
				    if(response.data=="pwdNA")

				    	alert("Password wrong")
				    else if(response.data=="usrNA")							    	
				    	alert("Email wrong")
				    else {
				    	document.cookie="Spanzhash="+response.data;
				    	$window.location.href = '/dashboard';
					}
				  }, function errorCallback(response) {				    			    	
				    	$scope.message.error="Server Error"
				  });
			} else alert("Introduceti si parola");
		} else alert("Introduceti email");
    }
    
});