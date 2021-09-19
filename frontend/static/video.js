var Main = {
    data() {
        return {
            canTouche: false,
            wsavc: null,
        };
    },
    mounted: function () {
        this.$nextTick(function () {
            this.canTouche = this.supportsTouches();
            if (!this.canTouche) {
                this.mountedKeyupEvent();
            }
            this.videoStreaming(getCookieValueByKey("token"));
        })
    },
    methods: {
        supportsTouches() {
            return ("createTouch" in document);
        },
        mountedKeyupEvent() {
            document.addEventListener('keydown', (e) => {
                var c = e.which || e.keyCode;
                switch (c) {
                    case 37:
                        // console.log("left");
                        this.sendServoCmd("left");
                        break
                    case 38:
                        // console.log("up");
                        this.sendServoCmd("up");
                        break
                    case 39:
                        // console.log("right");
                        this.sendServoCmd("right");
                        break
                    case 40:
                        // console.log("down")
                        this.sendServoCmd("down");
                        break
                }
            });
        },
        sendServoCmd(mod, value) {
            if (value == undefined || value == null || value <= 0)
                value = 0;
            this.wsavc.wsSend(JSON.stringify({ direction: mod, value: value }));
        },
        videoStreaming(token) {
            if (token == "" || token == undefined || token == null) {
                this.$message.error("authentication failed");
            }
            var canvas = document.getElementById("video");
            var uri = "ws://" + document.location.host + "/vi/stream?token=" + token;
            // Create h264 player
            var wsavc = new WSAvcPlayer(canvas, "webgl", 1, 35);
            wsavc.cmd({
                "action": "init",
                "width": 960,
                "height": 540
            });
            this.wsavc = wsavc;
            wsavc.connect(uri);
        }
    }
}
var Ctor = Vue.extend(Main)
new Ctor().$mount('#app')