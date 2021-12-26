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

	guangdongMapData = map[string]float64{
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
	fujiansheng = map[string]float64{
		"厦门市": float64(0),
		"莆田市": float64(1),
		"福州市": float64(2),
		"南平市": float64(3),
		"宁德市": float64(4),
		"三明市": float64(5),
		"泉州市": float64(6),
		"龙岩市": float64(7),
		"漳州市": float64(8),
	}
)