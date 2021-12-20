package maker

import (
	"github.com/varz1/nCovBot/channel"
	"strconv"
	"strings"
	"time"
)

func Overall() {
	text := strings.Builder{}
	for overall := range channel.OverallMsgChannel {
		data := overall.OverallData
		global := overall.OverallData.GlobalStatistics
		timeTem := "2006-01-02 15:04:05"
		tm := time.Unix(data.UpdateTime/1000, 0).Format(timeTem)
		text.WriteString("🇨🇳国内疫情概况:")
		text.WriteString("\n现存确诊(含港澳台):" + strconv.Itoa(data.CurrentConfirmedCount) + " ⬆️" + strconv.Itoa(data.CurrentConfirmedIncr))
		text.WriteString("\n现存无症状:" + strconv.Itoa(data.SeriousCount) + " ⬆️" + strconv.Itoa(data.SeriousIncr))
		text.WriteString("\n境外输入:" + strconv.Itoa(data.SuspectedCount) + " ⬆️" + strconv.Itoa(data.SuspectedIncr))
		text.WriteString("\n累计确诊:" + strconv.Itoa(data.ConfirmedCount) + " ⬆️" + strconv.Itoa(data.ConfirmedIncr))
		text.WriteString("\n累计治愈:" + strconv.Itoa(data.CuredCount) + " ⬆️" + strconv.Itoa(data.CuredIncr))
		text.WriteString("\n累计死亡" + strconv.Itoa(data.DeadCount) + " ⬆️" + strconv.Itoa(data.DeadIncr))
		text.WriteString("\n🌏全球疫情概况")
		text.WriteString("\n全球现存确诊" + strconv.Itoa(global.CurrentConfirmedCount) + " ⬆️" + strconv.Itoa(global.CurrentConfirmedIncr))
		text.WriteString("\n全球累计确诊" + strconv.Itoa(global.ConfirmedCount) + " ⬆️" + strconv.Itoa(global.ConfirmedIncr))
		text.WriteString("\n全球累计治愈" + strconv.Itoa(global.CuredCount) + " ⬆️" + strconv.Itoa(global.CuredIncr))
		text.WriteString("\n全球累计死亡" + strconv.Itoa(global.DeadCount) + " ⬆️" + strconv.Itoa(global.DeadIncr))
		text.WriteString("\n更新时间:" + tm)
		overall.Overall.Text = text.String()
		channel.MessageChannel <- overall.Overall
		text.Reset()
	}
}
