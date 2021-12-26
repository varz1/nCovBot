package maker

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"math/rand"
)

var (
	baseMapData = []opts.MapData{
		{Name: "北京", Value: float64(rand.Intn(150))},
		{Name: "上海", Value: float64(rand.Intn(150))},
		{Name: "广东", Value: float64(rand.Intn(150))},
		{Name: "辽宁", Value: float64(rand.Intn(150))},
		{Name: "山东", Value: float64(rand.Intn(150))},
		{Name: "山西", Value: float64(rand.Intn(150))},
		{Name: "陕西", Value: float64(rand.Intn(150))},
		{Name: "新疆", Value: float64(rand.Intn(150))},
		{Name: "内蒙古", Value: float64(rand.Intn(150))},
	}
	fujian = map[string]float32{
		"厦门市": float32(rand.Intn(150)),
		"福州市": float32(rand.Intn(150)),
		"莆田市": float32(rand.Intn(150)),
		"龙岩市": float32(rand.Intn(150)),
		"三明市": float32(rand.Intn(150)),
		"宁德市": float32(rand.Intn(150)),
		"泉州市": float32(rand.Intn(150)),
		"漳州市": float32(rand.Intn(150)),
		"南平市": float32(rand.Intn(150)),
	}
	guangdong = map[string]float64{
		"深圳市": float64(rand.Intn(150)),
		"广州市": float64(rand.Intn(150)),
		"湛江市": float64(rand.Intn(150)),
		"汕头市": float64(rand.Intn(150)),
		"东莞市": float64(rand.Intn(150)),
		"佛山市": float64(rand.Intn(150)),
		"云浮市": float64(rand.Intn(150)),
		"肇庆市": float64(rand.Intn(150)),
		"梅州市": float64(rand.Intn(150)),
	}
)
