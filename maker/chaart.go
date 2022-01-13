package maker

import (
	"bytes"
	data2 "github.com/varz1/nCovBot/data"
	"github.com/vdobler/chart"
	"github.com/vdobler/chart/imgg"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"strconv"
	"strings"
	"time"
)

type Dumper struct {
	N, M, W, H, Cnt int
	I               *image.RGBA
	img             bytes.Buffer
}

func NewDumper(n, m, w, h int) *Dumper {
	dumper := Dumper{N: n, M: m, W: w, H: h}
	dumper.I = image.NewRGBA(image.Rect(0, 0, n*w, m*h))
	bg := image.NewUniform(color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff})
	draw.Draw(dumper.I, dumper.I.Bounds(), bg, image.ZP, draw.Src)
	return &dumper
}

func (d *Dumper) Plot(c chart.Chart) error {
	row, col := d.Cnt/d.N, d.Cnt%d.N
	igr := imgg.AddTo(d.I, col*d.W, row*d.H, d.W, d.H, color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}, nil, nil)
	c.Plot(igr)
	err := png.Encode(&d.img, d.I)
	if err != nil {
		log.Println("编码PNG失败")
		return err
	}
	return nil
}

// Scatter 渲染散点图
func Scatter(x, y []float64, chartName string) *bytes.Buffer {
	dumper := NewDumper(1, 1, 800, 600)
	const DAY float64 = 86400 // 一天的时间戳
	pl := chart.ScatterChart{Title: chartName}
	pl.Key.Pos = "itl"
	todayAdd := "Today Add " + strconv.FormatFloat(y[len(y)-1], 'g', -1, 32) //将今日新增标注上
	pl.AddDataPair(todayAdd, x, y, chart.PlotStyleLinesPoints,
		chart.Style{Symbol: '#', SymbolColor: color.NRGBA{R: 0xE3, G: 0x17, B: 0x0D, A: 0xff}, LineStyle: chart.SolidLine})
	pl.XRange.TicSetting.Mirror = 1
	pl.XRange.TicSetting.TLocation = time.Local
	pl.XRange.Time = true
	pl.XRange.DataMin = x[0]
	pl.XRange.DataMax = x[len(x)-1] - DAY
	pl.XRange.TicSetting.TDelta = chart.MatchingTimeDelta(float64(time.Now().Unix()), DAY) //x轴时间间隔

	pl.YRange.TicSetting.Mirror = 1
	pl.XRange.Label = "date"
	pl.YRange.Label = "cases"
	err := dumper.Plot(&pl)
	if err != nil {
		return nil
	}
	return &dumper.img
}

// PieChart 渲染饼状图
func PieChart(continent map[string]int, chartName string) *bytes.Buffer {
	dumper := NewDumper(1, 1, 500, 300)

	var names []string
	var cases []int
	for k, v := range continent {
		if k == "PubDate" {
			continue
		}
		names = append(names, k)
		cases = append(cases, v)
	}

	pie := chart.PieChart{Title: chartName}
	pie.AddIntDataPair("World", names, cases)
	pie.Data[0].Samples[3].Flag = true

	pie.Inner = 0.55 //面积比例
	pie.FmtVal = chart.PercentValue
	err := dumper.Plot(&pie)
	if err != nil {
		return nil
	}
	return &dumper.img
}

func GetScatter() error {
	log.Println("开始绘图Trend")
	var err error
	const Day = 86400
	adds := data2.GetAdds(7) //获取七天本地新增
	if adds == nil {
		return err
	}
	var xRange, yRange []float64
	for _, v := range adds {
		s := strings.ReplaceAll(v.Date, ".", "")
		res := Time2TimeStamp(s)
		xRange = append(xRange, float64(res+Day))
		yRange = append(yRange, float64(v.LocalConfirmAdd))
	}
	buf := Scatter(xRange, yRange, "Local Cases Increment In 7 Days")
	if buf == nil {
		return err
	}
	SCATTER = *buf
	return nil
}
