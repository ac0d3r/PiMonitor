var Main = {
    data() {
        return {
            userLoginStatus: false, // false 未登录
            logoSrc: './static/logo.png',
            loginForm: {
                name: '',
                password: '',
            },
            updateFormVisible: false,
            updateForm: {
                name: '',
                password: '',
            },
        };
    },
    mounted: function () {
        document.addEventListener('keydown', (e) => {
            var c = e.which || e.keyCode;
            switch (c) {
                case 13: // 回车登陆
                    this.userLogin();
                    break
            }
        });
        this.$nextTick(function () {
            let token = getCookieValueByKey("token")
            if (token == null) {
                this.userLoginStatus = false;
            } else {
                this.userLoginStatus = true;
            }
        })
    },
    methods: {
        userLogin() {
            let that = this;
            var http = getHTTPObject();
            if (http == null) {
                return
            }
            http.onload = function (e) {
                let data = JSON.parse(http.response);
                if (data.code > 0) {
                    that.parseErrorResponse(data);
                } else {
                    saveUserToken(data.data);
                    that.userLoginStatus = true;
                }
            }
            http.open("POST", getUrl("/login"));
            sendJSON(http, { username: this.loginForm.name, password: this.loginForm.password })
        },
        userLogout() {
            let that = this;
            var http = getHTTPObject();
            if (http == null) {
                return
            }
            http.onload = function (e) {
                removeUserToken();
                let data = JSON.parse(http.response);
                if (data.code > 0) {
                    that.parseErrorResponse(data);
                } else {
                    removeUserToken();
                    that.userLoginStatus = false;
                }
            }
            http.open("POST", getUrl("/api/user/logout"));
            carryUserToken(http);
            http.send();
        },
        userUpdateInfomation() {
            this.updateFormVisible = false;
            let that = this;
            var http = getHTTPObject();
            if (http == null) {
                return
            }
            http.onload = function (e) {
                let data = JSON.parse(http.response);
                if (data.code > 0) {
                    that.parseErrorResponse(data);
                } else {
                    that.$message.success("update user infomation success");
                }
            }
            http.open("POST", getUrl("/api/user/update"));
            carryUserToken(http);
            sendJSON(http, { username: this.updateForm.name, password: this.updateForm.password })
        },
        parseErrorResponse(data) {
            this.$message.error(data.msg);
            switch (data.code) {
                case 2000:
                case 2001:
                case 2002:
                    this.userLoginStatus = false;
                    removeUserToken();
                    break;
            }
        },
        goVideo() {
            window.open("http://" + document.location.host + '/video.html');
        }
    }
}
var Ctor = Vue.extend(Main)
new Ctor().$mount('#app')
