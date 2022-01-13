package maker

import (
	"bytes"
	"github.com/vdobler/chart"
	"github.com/vdobler/chart/imgg"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
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
	var STEP float64 = 86400 // 一天的时间戳
	pl := chart.ScatterChart{Title: chartName}
	pl.Key.Pos = "itl"
	pl.AddDataPair("case", x, y, chart.PlotStyleLinesPoints,
		chart.Style{Symbol: '#', SymbolColor: color.NRGBA{R: 0xE3, G: 0x17, B: 0x0D, A: 0xff}, LineStyle: chart.SolidLine})
	pl.XRange.TicSetting.Mirror = 1
	pl.XRange.TicSetting.TLocation = time.Local
	pl.XRange.Time = true
	pl.XRange.DataMin = x[0]
	pl.XRange.DataMax = x[len(x)-1]-STEP
	pl.XRange.TicSetting.TDelta = chart.MatchingTimeDelta(float64(time.Now().Unix()), STEP) //x轴时间间隔

	pl.YRange.TicSetting.Mirror = 1
	pl.XRange.Label = "date"
	pl.YRange.Label = "cases"
	err := dumper.Plot(&pl)
	if err != nil {
		return nil
	}
	return &dumper.img
}
