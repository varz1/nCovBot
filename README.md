# nCovBot

## TODO

- [x] 全局概览
- [x] 优化图表
- [ ] 优化日志
- [ ] 定时任务错误处理


## 简介

一个Telegram Covid-19 Bot

![nCovBot](https://github.com/varz1/pics/blob/master/bot.png?raw=true)

## ChangeLog

- **2022.3.11** 添加环境变量初始化,设置Http路由,使用Rod替换chromeDp
- **2022.2.21** 优化model
- **2022.2.08** 实现了一个内存kv存储器,将上游的信息进行缓存并定期同步,提高服务响应速度和稳定性
- **2022.1.20** 全局概览
- **2022.1.16** 初始化数据,定时更新数据,风险地区优化
- **2022.1.15** 图表中文支持
- **2022.1.14** 使用文件传输图表,优化了定时器,初始化图表

## 数据以及图表来源

[丁香园](https://github.com/BlankerL/DXY-COVID-19-Data)
[腾讯](https://news.qq.com/zt2020/page/feiyan.htm#/)
[百度](https://voice.baidu.com/act/newpneumonia/newpneumonia)


## 部分功能

### 折线图

![1](https://github.com/varz1/pics/blob/master/Snipaste_2022-11-14_16-28-14.png?raw=true)

### 饼图
![2](https://github.com/varz1/pics/blob/master/Snipaste_2022-11-14_16-28-26.png?raw=true)


### 风险地区
![3](https://github.com/varz1/pics/blob/master/Snipaste_2022-11-14_16-29-39.png?raw=true)


### 概览
![4](https://github.com/varz1/pics/blob/master/Snipaste_2022-11-14_16-30-06.png?raw=true)


### 实时新闻
![5](https://github.com/varz1/pics/blob/master/Snipaste_2022-11-14_16-30-52.png?raw=true)