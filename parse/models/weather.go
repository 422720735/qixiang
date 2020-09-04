package models

type CityLive struct {
	Today           string
	Update          string
	Date            string
	SubTitle        string
	Week            string
	Temperature     string // 气温
	Situation       string // 多云转晴
	AdviceImg       string // 当天图片
	BodyTemperature string // 体温
	Wet             string // 湿度
	Press           string // 气压
	Visibility      string // 能见度
	Rainfall        string // 降雨
	Wind            string // 风
	Days3           string // 近3天
	Exponent        string // 生活指数
}

type Pm25 struct {
	CityName   string // 城市
	AirQuality string // 空气质量
	Health     string // 对健康影响情况： 空气质量可接受，但某些污染物可能对极少数异常、敏感人群健康有较弱影响。
	Advice     string // 建议采取的措施： 除少数对某些污染物特别容易过敏的人群外，其他人群可以正常进行室外活动。
}

// 宠物知识点
type Knowledge struct {
	ImageList string // 搞笑萌图

	Title   string
	Content string
}

type Weather struct {
	*CityLive
	*Pm25
	*Knowledge
}

var Data *Weather
