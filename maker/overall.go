package maker

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/robfig/cron/v3"
	"github.com/varz1/nCovBot/cache"
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
	timer = cron.New()
	C     = cache.New()
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
		oa, exist := data2.C.Get("overall")
		m, _ := C.Get("map")
		Map := m.(model.Chartt)
		if Map.Pie.Bytes() == nil || !exist {
			errMsg := tgbotapi.NewMessage(overall.Message.Chat.ID, "获取数据失败")
			channel.MessageChannel <- errMsg
			return
		}
		data1 := oa.(model.Overall)
		data := data1.Data.Diseaseh5Shelf
		tm :=data.LastUpdateTime
		total := data.ChinaTotal
		add := data.ChinaAdd
		caption.WriteString("🇨🇳中国疫情概况:")
		caption.WriteString("\n本土现有确诊:" + strconv.Itoa(total.LocalConfirm) + " 较上日⬆️" + strconv.Itoa(add.LocalConfirmH5))
		caption.WriteString("\n现存确诊(含港澳台):" + strconv.Itoa(total.NowConfirm) + " 较上日⬆️" + strconv.Itoa(add.NowConfirm))
		caption.WriteString("\n累计确诊:" + strconv.Itoa(total.Confirm) + " 较上日⬆️" + strconv.Itoa(add.Confirm))
		caption.WriteString("\n无症状感染者:" + strconv.Itoa(total.NoInfect) + " 较上日⬆️" + strconv.Itoa(add.NoInfect))
		caption.WriteString("\n境外输入:" + strconv.Itoa(total.ImportedCase) + " 较上日⬆️" + strconv.Itoa(add.ImportedCase))
		//caption.WriteString("\n累计治愈:" + strconv.Itoa(data.CuredCount) + " 较上日⬆️" + strconv.Itoa(data.CuredIncr))
		caption.WriteString("\n累计死亡" + strconv.Itoa(total.Dead) + " 较上日⬆️" + strconv.Itoa(add.Dead))
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
		s, _ := C.Get("scatter")
		SCATTER := s.(model.Chartt)
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
		world, exist := data2.C.Get("world")
		data := world.(model.OverallWorld)
		p1, _ := C.Get("pie")
		Pie := p1.(model.Chartt)
		global := data.Data.WomWorld
		if Pie.Pie.Bytes() == nil || !exist {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "数据错误")
			channel.MessageChannel <- msg
			return
		}
		caption := strings.Builder{}
		tm := global.LastUpdateTime
		caption.WriteString("\n🌏全球疫情概况")
		caption.WriteString("\n全球现存确诊" + strconv.Itoa(global.NowConfirm) + " 较上日⬆️" + strconv.Itoa(global.NowConfirmAdd))
		caption.WriteString("\n全球累计确诊" + strconv.Itoa(global.Confirm) + " 较上日⬆️" + strconv.Itoa(global.ConfirmAdd))
		caption.WriteString("\n全球累计治愈" + strconv.Itoa(global.Heal) + " 较上日⬆️" + strconv.Itoa(global.HealAdd))
		caption.WriteString("\n全球累计死亡" + strconv.Itoa(global.Dead) + " 较上日⬆️" + strconv.Itoa(global.DeadAdd))
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
