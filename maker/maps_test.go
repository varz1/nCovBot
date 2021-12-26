package maker

import (
	"github.com/go-resty/resty/v2"
	"io"
	"os"
	"testing"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var request =resty.New()
//func generateMapData(data map[string]float64) (items []opts.MapData) {
//	items = make([]opts.MapData, 0)
//	for k, v := range data {
//		items = append(items, opts.MapData{Name: k, Value: v})
//	}
//	return
//}

func mapVisualMap() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")
	mc.SetGlobalOptions(
		charts.WithToolboxOpts(opts.Toolbox{
			Show: true,
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show: true, Type: "png",
				}}}),
		charts.WithTitleOpts(opts.Title{
			Title: "VisualMap",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
		}),
	)
	mc.AddSeries("map", baseMapData).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)
	return mc
}

//func mapRegion() *charts.Map {
//	mc := charts.NewMap()
//	mc.RegisterMapType("福建")
//	mc.SetGlobalOptions(
//		charts.WithTitleOpts(opts.Title{
//			Title: "福建省",
//		}),
//		charts.WithVisualMapOpts(opts.VisualMap{
//			Calculable: true,
//		}),
//	)
//	mc.AddSeries("map", generateMapData(fujiansheng)).SetSeriesOptions(
//		charts.WithLabelOpts(opts.Label{
//			Show: true,
//		}))
//	return mc
//}

type MapExamples struct{}

func (MapExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		mapVisualMap(),
	)
	f, err := os.Create("map.html")
	if err != nil {
		panic(err)
	}
	err1 := page.Render(io.MultiWriter(f))
	if err1 != nil {
		return
	}
}

func TestGetMap(t *testing.T) {
	//resp,err := request.R().SetResult().SetQueryString("modules=provinceCompare").
	//	Get("https://api.inews.qq.com/newsqa/v1/query/inner/publish/modules/list?")
	ex := new(MapExamples)
	ex.Examples()
}

type provinceCompare struct {

}