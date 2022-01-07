# nCovBot

## TODO

- [ ] 疫情图动态标签需要更新
- [ ] 优化日志

## 简介

一个Telegram的聊天机器人疫情信息聚合机器人,数据来自丁香园开放Api，图表来自百度，使用无头浏览器爬取。

![nCovBot](https://github.com/varz1/pics/blob/master/bot.png?raw=true)

## 第三方模块

- Http客户端 — [Resty](https://github.com/go-resty/resty)
- 消息 — [Telegram bot](https://github.com/go-telegram-bot-api/telegram-bot-api)
- 日志 — [Logrus](https://github.com/spf13/viper)
- 无头浏览器—[ChromeDp](https://github.com/chromedp/chromedp)
- net框架—[Fiber](https://github.com/gofiber/fiber)

## 主要功能

- 支持地区列表
- 国内疫情概览（包括中国疫情地图
- 本土疫情趋势图
- 最新新闻（返回新闻概览以及列表
- 中高风险地区查询
- 输入地区名称返回疫情数据

## 我的环境以及用到的服务
- Go 1.17
- Debian 10
- Telegram [Api](https://github.com/go-telegram-bot-api/telegram-bot-api)
- [Heroku](https://dashboard.heroku.com/apps)
- Telegram 客户端


