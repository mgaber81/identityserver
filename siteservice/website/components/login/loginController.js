(function () {
    'use strict';
    angular.module('loginApp')
        .controller('loginController', ['$http', '$window', '$scope', '$rootScope', '$interval', 'LoginService', loginController]);

    function loginController($http, $window, $scope, $rootScope, $interval, LoginService) {
        var vm = this;
        var urlParams = URI($window.location.href).search(true);
        vm.submit = submit;
        vm.clearValidation = clearValidation;
        vm.validateUsername = validateUsername;
        vm.resetValidation = resetValidation;
        vm.loginInfoValid = loginInfoValid;
        vm.externalSite = urlParams.client_id;
        $rootScope.registrationUrl = '/register' + $window.location.search;
        vm.logo = "";
        vm.twoFAMethod = 'sms';
        vm.login = "";
        vm.password = "";
        vm.description = ""

        var listener;
        activate();

        function activate() {
            requestRegister();
            if (vm.externalSite) {
                LoginService.getLogo(vm.externalSite).then(
                    function(data) {
                        vm.logo = data.logo;
                        resizeLogo();
                    }
                );
                $window.addEventListener('resize', resizeLogo, false);
                $window.addEventListener('orientationchange', resizeLogo, false);
                loadDescription();
            }
            autoFillListener();
            $scope.$on('$destroy', function() {
                  // Make sure that the interval is destroyed too
                  stopListening();
            });
        }

        // Load the correct description after the user changes language
        $rootScope.$on('$translateChangeSuccess', function () {
            loadDescription();
        });

        function loadDescription() {
            LoginService.getDescription(vm.externalSite, localStorage.getItem('langKey')).then(
                function(data) {
                    vm.description = data.text;
                }
            );
        }

        function renderLogo() {
            if (vm.logo !== "") {
                var img = new Image();
                img.onload = function() {
                    var c = document.getElementById("login-logo");
                    if (!c) {
                        console.log("aborting logo render - canvas not loaded");
                        return;
                    }
                    var ctx = c.getContext("2d");
                    ctx.clearRect(0, 0, c.width, c.height);
                    ctx.drawImage(img, 0, 0, c.width, c.height);
                };
                img.src = vm.logo;

            }
        }

        function autoFillListener() {
            listener = $interval(function() {
                var login = document.getElementById("username");
                var password = document.getElementById("password");
                if (login.value !== vm.login) {
                    vm.login = login.value;
                }
                if (password.value !== vm.password) {
                    vm.password = password.value;
                }
            }, 100);
        }

        function stopListening() {
            if (angular.isDefined(listener)) {
                $interval.cancel(listener);
                listener = undefined;
            }
        }

        function submit() {
            var data = {
                login: vm.login,
                password: vm.password
            };
            var url = '/login' + $window.location.search;
            $http.post(url, data).then(
                function (data) {
                  if (data.data.redirecturl) {
                      // Skip 2FA when logging in from an external site if the 2FA validity period hasn't passed
                      $window.location.href = data.data.redirecturl;
                  } else {
                      // Redirect 2 factor authentication page
                      $window.location.hash = '#/2fa';
                  }
                },
                function (response) {
                    if (response.status === 422) {
                        $scope.loginform.password.$setValidity("invalidcredentials", false);
                    }
                }
            );
        }

        function clearValidation() {
            $scope.loginform.password.$setValidity("invalidcredentials", true);
        }

        function validateUsername(username) {
            var options = {
                params: {
                    username: username
                }
            };
            return $http.get('/validateusername', options);
        }

        function resetValidation(prop) {
            switch (prop) {
                case 'phonenumber':
                    $scope.loginform[prop].$setValidity("invalid_phonenumber", true);
                    break;
                case 'totpcode':
                    $scope.loginform[prop].$setValidity("invalid_totpcode", true);
                    break;
                case 'twoFAMethod':
                    $scope.loginform.phonenumber.$setValidity("invalid_phonenumber", true);
                    $scope.loginform.phonenumber.$setValidity("pattern", true);
                    $scope.loginform.totpcode.$setValidity("totpcode", true);
                    break;
            }
        }

        function loginInfoValid() {
            return $scope.loginform.username
                && $scope.loginform.username.$valid
                && $scope.loginform.password.$valid;
        }

        function signupInfoValid() {
            switch (vm.twoFAMethod) {
                case 'sms':
                    return basicInfoValid() && $scope.loginform.phonenumber.$valid;
                    break;
                case 'totp':
                    return basicInfoValid() && $scope.loginform.totpcode.$valid;
                    break;
            }
        }

        function resizeLogo(e) {
            var formArea = document.getElementById("form-area");
            var logoArea = document.getElementById("login-logo");
            var widthToHeight = 25 / 12;
            var newWidth = formArea.clientWidth - 20;
            if (newWidth < 500) {
                logoArea.width = newWidth;
                logoArea.height = (newWidth) / widthToHeight;
            } else if (newWidth >= 500 && logoArea.width < 500) {
                logoArea.width = 500;
                logoArea.height = 240;
            }
            renderLogo();
        }

        // Check if the 'prefer' queryparam is set to register. If so, and there has been no previous login on this device
        // according to the localStorage, remove the prefer param and redirect to the register page.
        function requestRegister() {
            if (localStorage.getItem('hasLoggedIn')) {
                return;
            }
            if (urlParams.prefer === "register") {
                delete urlParams.prefer;
                var url = "/register?";
                // Manually reconstruct the URL because angular doesn't seem to be able to delete just 1 queryparam at the moment.
                angular.forEach(urlParams, function(value, key) {
                    url = url + key + "=" + value + "&";
                });
                url = url.slice(0, -1);
                $window.location.href = url;
            }
        }
    }
})();
