<div align="center" >
    <img src="./frontend/static/logo.png" width="280" alt="PiMonitor" />
    <h1>PiMonitor</h1>
</div>

![](https://img.shields.io/badge/Golang-1.16+-292e33?style=flat-square&logo=go) + ![](https://img.shields.io/badge/RaspberryPi-4b-292e33?style=flat-square&logo=raspberry-pi) 实现可控制的WEB视频监控器。

---

## Hardware

- 树莓派4b
- 舵机(SG90)*2
- 摄像头(Raspberry Pi Camera)
- 杜邦线若干

**接线：**
- 摄像头接线(goooooogle
- 舵机我分为了X轴和Y轴，X轴信号线：GPIO-23，Y轴信号线：GPIO-18；

可随意修改，注意同步 `pimonitor.go` 代码中舵机设置：
```golang
...
// 初始化舵机
servo.Init("16", "12", 15, 35, 10, 40)
...
```

贴心附上GPIO图:
![GPIO-Pinout-Diagram](https://user-images.githubusercontent.com/26270009/133922602-46dbe000-26df-491b-b8c7-09f9f7e6df63.png)


## How to run the server

1. 【Enabling PWM output on GPIO pins】https://gobot.io/documentation/platforms/raspi/

（尽管如此也不能完美的控制舵机，不过因摄像头排线的问题只能在一个合理范围角度内转动，这目前测试下来还是没有出什么问题的。

2. 编译 && 启动服务
```bash
$ git clone https://github.com/Buzz2d0/PiMonitor.git
$ cd PiMonitor
$ make init
```

# Usage

**默认用户名密码：** `admin`:`admin`

用浏览器访问 http://[your-raspberry-pi-ip]:8080 登陆后，可以查看实时视频画面（PC端下使用方向键控制舵机）。

## TODO

- [ ] 触摸屏下虚拟按键控制舵机