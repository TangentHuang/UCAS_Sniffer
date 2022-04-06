# 网络攻防基础实验—网络嗅探器

国科大 网络攻防基础 2021-2022 春季 课程作业

## 项目介绍

github仓库：https://github.com/TangentHuang/UCAS_Sniffer
本项目使用golang语言开发，gui使用gin框架，核心功能基于gopacket库开发。仅在win11下测试，基于go语言跨平台特性，linux和macos环境理论上均可以运行。

```
.
├── MySniffer.go
├── capture.go
├── go.mod
├── go.sum
├── main.go
└── makeMainContent.go
```

## 项目安装

本项目基于goland IED开发，也可以直接用命令行编译运行，具体如下。

安装流程

```shell
git clone https://github.com/TangentHuang/UCAS_Sniffer.git
go mod init ucas_sniffer
go get -d -v ./...
go build 
./ucas_sniffer.exe
```

## 项目功能

提供了网络嗅探抓包的功能，支持tcp、udp、arp、icmp等多种协议解析。提供是网卡选择功能，可以进行实时抓取网卡上的数据包，同时提供了解析pcap文件和保存抓包内存为pcap文件的功能。提供了基于BPF格式的包过滤功能。

### 主界面

![](https://img.tangent.ink/20220406214559.png)

最上面为菜单栏，提供了打开和保存pcap文件的选项，interfaceMenu下可以选择网卡。工具栏提供了开始和停止抓包的按钮，当前选择网卡的名字和BPF过滤器的输入框。

主界面从上到下分别显示packet list、packet details、packet in binary。

### 打开和保存pcap文件

在菜单栏file中，选择open，即可打开pcap文件

<img src="https://img.tangent.ink/20220406164240.png" style="zoom:50%;" />

在菜单栏file中，选择save，即可把抓取到的包保存为pcap文件

<img src="https://img.tangent.ink/20220406164437.png" style="zoom:50%;" />

### 网卡选择

在菜单栏interfaceMenu中，选择interface，即可弹出网卡选择界面选择网卡。

![](https://img.tangent.ink/20220406164326.png)

### 实时抓包

在选择好网卡过后，点击工具栏中的开始按钮，即可开始抓包，点击停止按钮，即可停止抓包

### BPF包过滤

在工具栏BPF输入框中输入BPF指令，点击submit按钮，即可实现包过滤
