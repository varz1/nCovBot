# nCovBot

## TODO

- [ ] 优化日志
- [ ] 优化图表

## 简介

一个Telegram的聊天机器人疫情信息聚合机器人,数据来自丁香园开放Api，地图来自百度，使用无头浏览器爬取,本土疫情趋势后端绘制数据来自百度。

![nCovBot](https://github.com/varz1/pics/blob/master/bot.png?raw=true)

## 主要功能

- 支持地区列表
- 国内疫情概览（包括中国疫情地图
- 本土疫情趋势图
- 最新新闻（返回新闻概览以及列表
- 中高风险地区查询
- 输入地区名称返回疫情数据

## 数据以及图表来源

[丁香园](https://github.com/BlankerL/DXY-COVID-19-Data)
[腾讯](https://news.qq.com/zt2020/page/feiyan.htm#/)
[百度](https://voice.baidu.com/act/newpneumonia/newpneumonia)

## ChangeLog
**2022.1.14** V95:使用文件传输图表,优化了定时器，初始化图表