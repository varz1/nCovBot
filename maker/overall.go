package maker

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/robfig/cron/v3"
	"github.com/varz1/nCovBot/channel"
	data2 "github.com/varz1/nCovBot/data"
	"github.com/varz1/nCovBot/model"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	timer   = cron.New()
	SCATTER = model.Chartt{}
	Pie     = model.Chartt{}
	Map     = model.Chartt{}
)

// 初始化图表以及定时任务
func init() {
	GetScatter()
	GetPie()
	GetChMap()
	timer.AddFunc("@every 6h", func() {
		GetScatter()
		GetChMap()
		GetPie()
	})
	timer.AddFunc("@every 30m", func() {
		resp, err := http.Get("https://ncovbott.herokuapp.com/hi")
		if err != nil {
			log.Println("定时ping失败")
			return
		}
		log.Printf("Ping成功 %v", resp.StatusCode)
	})
	timer.Start()
}

func Overall() {
	caption := strings.Builder{}
	for overall := range channel.OverallUpdateChannel {
		if Map.Pie.Bytes() == nil {
			errMsg := tgbotapi.NewMessage(overall.Message.Chat.ID, "获取图表失败")
			channel.MessageChannel <- errMsg
			return
		}
		oa, _ := data2.C.Get("overall")
		data := oa.(model.OverallData)
		tm := time.Unix(data.UpdateTime/1000, 0).Format("2006-01-02 15:04")
		caption.WriteString("🇨🇳中国疫情概况:")
		caption.WriteString("\n现存确诊(含港澳台):" + strconv.Itoa(data.CurrentConfirmedCount) + " ⬆️" + strconv.Itoa(data.CurrentConfirmedIncr))
		caption.WriteString("\n现存无症状:" + strconv.Itoa(data.SeriousCount) + " ⬆️" + strconv.Itoa(data.SeriousIncr))
		caption.WriteString("\n境外输入:" + strconv.Itoa(data.SuspectedCount) + " ⬆️" + strconv.Itoa(data.SuspectedIncr))
		caption.WriteString("\n累计确诊:" + strconv.Itoa(data.ConfirmedCount) + " ⬆️" + strconv.Itoa(data.ConfirmedIncr))
		caption.WriteString("\n累计治愈:" + strconv.Itoa(data.CuredCount) + " ⬆️" + strconv.Itoa(data.CuredIncr))
		caption.WriteString("\n累计死亡" + strconv.Itoa(data.DeadCount) + " ⬆️" + strconv.Itoa(data.DeadIncr))
		caption.WriteString("\n数据更新时间:" + tm)
		msg := tgbotapi.PhotoConfig{
			BaseFile: tgbotapi.BaseFile{
				BaseChat: tgbotapi.BaseChat{ChatID: overall.Message.Chat.ID},
				File: tgbotapi.FileBytes{
					Name:  "map.jpg",
					Bytes: Map.Pie.Bytes(),
				},
			},
			Caption: caption.String() + "\n图表更新时间:" + Map.Date,
		}
		channel.MessageChannel <- msg
		caption.Reset()
	}
}

func Trend() {
	for update := range channel.TrendChannel {
		if SCATTER.Pie.Bytes() == nil {
			errMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "渲染错误")
			channel.MessageChannel <- errMsg
			return
		}
		msg := tgbotapi.PhotoConfig{
			BaseFile: tgbotapi.BaseFile{
				BaseChat: tgbotapi.BaseChat{ChatID: update.Message.Chat.ID},
				File: tgbotapi.FileBytes{
					Name:  "trend.jpg",
					Bytes: SCATTER.Pie.Bytes(),
				},
			},
			Caption: "图表更新时间:" + SCATTER.Date,
		}
		channel.MessageChannel <- msg
	}
}

func WorldOverall() {
	for update := range channel.WorldUpdateChannel {
		data := data2.OA
		global := data.GlobalStatistics
		if Pie.Pie.Bytes() == nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "渲染错误")
			channel.MessageChannel <- msg
			return
		}
		caption := strings.Builder{}
		tm := time.Unix(data.UpdateTime/1000, 0).Format("2006-01-02 15:04")
		caption.WriteString("\n🌏全球疫情概况")
		caption.WriteString("\n全球现存确诊" + strconv.Itoa(global.CurrentConfirmedCount) + " ⬆️" + strconv.Itoa(global.CurrentConfirmedIncr))
		caption.WriteString("\n全球累计确诊" + strconv.Itoa(global.ConfirmedCount) + " ⬆️" + strconv.Itoa(global.ConfirmedIncr))
		caption.WriteString("\n全球累计治愈" + strconv.Itoa(global.CuredCount) + " ⬆️" + strconv.Itoa(global.CuredIncr))
		caption.WriteString("\n全球累计死亡" + strconv.Itoa(global.DeadCount) + " ⬆️" + strconv.Itoa(global.DeadIncr))
		caption.WriteString("\n数据更新时间:" + tm)
		caption.WriteString("\n图表更新时间:" + Pie.Date)
		p := tgbotapi.PhotoConfig{
			BaseFile: tgbotapi.BaseFile{
				BaseChat: tgbotapi.BaseChat{
					ChatID: update.Message.Chat.ID,
				},
				File: tgbotapi.FileBytes{
					Name:  "world.jpg",
					Bytes: Pie.Pie.Bytes(),
				},
			},
			Caption: caption.String(),
		}
		channel.MessageChannel <- p
		caption.Reset()
	}
}

func Time2TimeStamp(t string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt, _ := time.ParseInLocation("20060102", "2022"+t, loc)
	return tt.Unix()
}
