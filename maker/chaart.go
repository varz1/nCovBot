package maker

import (
	"bytes"
	"context"
	"github.com/chromedp/chromedp"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/sirupsen/logrus"
	data2 "github.com/varz1/nCovBot/data"
	"github.com/varz1/nCovBot/model"
	"github.com/varz1/nCovBot/variables"
	"github.com/vdobler/chart"
	"github.com/vdobler/chart/imgg"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var font *truetype.Font

// 初始化字体
func init() {
	var err error
	font, err = freetype.ParseFont(YaHeiFontData())
	if err != nil {
		panic(err)
	}
}

type Dumper struct {
	N, M, W, H, Cnt int
	I               *image.RGBA
	img             bytes.Buffer
}

func NewDumper(n, m, w, h int) *Dumper {
	dumper := Dumper{N: n, M: m, W: w, H: h}
	dumper.I = image.NewRGBA(image.Rect(0, 0, n*w, m*h))
	bg := image.NewUniform(color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff})
	draw.Draw(dumper.I, dumper.I.Bounds(), bg, image.Point{}, draw.Src)
	return &dumper
}

func (d *Dumper) Plot(c chart.Chart) error {
	row, col := d.Cnt/d.N, d.Cnt%d.N
	igr := imgg.AddTo(d.I, col*d.W, row*d.H, d.W, d.H, color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}, font, nil)
	c.Plot(igr)
	err := png.Encode(&d.img, d.I)
	if err != nil {
		logrus.Error("编码PNG失败")
		return err
	}
	return nil
}

// Scatter 渲染散点图
func Scatter(x, y []float64) *bytes.Buffer {
	dumper := NewDumper(1, 1, 800, 600)
	pl := chart.ScatterChart{Title: variables.TREND}
	pl.Key.Pos = "itl"
	todayAdd := "单位/例" //将今日新增标注上
	pl.AddDataPair(todayAdd, x, y, chart.PlotStyleLinesPoints,
		chart.Style{Symbol: '#', SymbolColor: color.NRGBA{R: 0xE3, G: 0x17, B: 0x0D, A: 0xff}, LineStyle: chart.SolidLine})
	pl.XRange.TicSetting.Mirror = 1
	pl.XRange.TicSetting.TLocation = time.Local
	pl.XRange.Time = true
	pl.XRange.DataMin = x[0]
	pl.XRange.DataMax = x[len(x)-1] - variables.DAY
	pl.XRange.TicSetting.TDelta = chart.MatchingTimeDelta(float64(time.Now().Unix()), variables.DAY) //x轴时间间隔

	pl.YRange.TicSetting.Mirror = 1
	pl.XRange.Label = "日期"
	pl.YRange.Label = "病例数"
	pl.YRange.Min = 0
	pl.YRange.Max = 300
	err := dumper.Plot(&pl)
	if err != nil {
		return nil
	}
	return &dumper.img
}

// PieChart 渲染饼状图
func PieChart(continent map[string]int) *bytes.Buffer {
	dumper := NewDumper(1, 1, 500, 300)
	var names []string
	var cases []int
	for k, v := range continent {
		names = append(names, k)
		cases = append(cases, v)
	}

	pie := chart.PieChart{Title: variables.PIE}
	pie.AddIntDataPair("大洲", names, cases)
	pie.Data[0].Samples[3].Flag = true

	pie.Inner = 0.55 //面积比例
	pie.FmtVal = chart.PercentValue
	err := dumper.Plot(&pie)
	if err != nil {
		return nil
	}
	return &dumper.img
}

// GetScatter 生成Scatter图
func GetScatter() {
	logrus.WithField("开始绘图Trend", "GetScatter")
	adds := data2.GetAdds(7) //获取七天本地新增
	if adds == nil {
		logrus.Error("数据为空")
	}
	var xRange, yRange []float64
	for _, v := range adds {
		s := strings.ReplaceAll(v.Date, ".", "")
		res := Time2TimeStamp(s)
		xRange = append(xRange, float64(res+86400))
		yRange = append(yRange, float64(v.LocalConfirmAdd))
	}
	buf := Scatter(xRange, yRange)
	if buf == nil {
		logrus.Error("渲染失败")
		return
	}
	var SCATTER model.Chartt
	SCATTER.Pie = *buf
	SCATTER.Date = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04")
	C.Set("scatter", SCATTER)
	logrus.Info("渲染成功")
}

// GetPie 生成饼状图
func GetPie() {
	logrus.WithField("开始绘制Pie", "GetPie")
	c, err1 := data2.GetWorldData()
	if err1 != nil {
		logrus.Error("获取Pie数据失败")
		return
	}
	buf := PieChart(c)
	var Pie model.Chartt
	Pie.Pie = *buf
	Pie.Date = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04")
	C.Set("pie", Pie)
	logrus.Info("渲染成功")
}

// GetChMap 无头浏览器爬取疫情地图
func GetChMap() {
	logrus.WithField("开始爬取地图", "GetChMap")
	var url = "https://voice.baidu.com/act/newpneumonia/newpneumonia"
	var selMap = "#virus-map"
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	chromedp.ExecPath(os.Getenv("GOOGLE_CHROME_SHIM"))
	// 超时设置
	ctx, cancel = context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()
	var buf []byte
	if err := chromedp.Run(ctx,
		Screenshot(url, selMap, &buf)); err != nil {
		logrus.Error("截图失败")
	}
	var Map model.Chartt
	Map.Pie.Write(buf)
	Map.Date = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04")
	C.Set("map",Map)
	logrus.Info("截图成功")
}

// Screenshot 截图
func Screenshot(url, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}

// YaHeiFontData 微软雅黑
func YaHeiFontData() []byte {
	p, _ := os.Getwd()
	fontBytes, err := ioutil.ReadFile(p + "/.fonts/WeiRuanYaHei-1.ttf")
	if err != nil {
		log.Println(err)
		return nil
	}
	return fontBytes
}
