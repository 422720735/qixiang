package weather

import (
	"github.com/PuerkitoBio/goquery"
	"qixiang/parse/models"
	"qixiang/parse/regular"
	"strings"
)

func Today(content []byte) (*models.CityLive, error) {
	var result models.CityLive
	dom, err := regular.NewDom(content)

	if err != nil {
		return nil, err
	}

	dom.Find("table.tableTop tbody tr td").Each(func(i int, item *goquery.Selection) {
		if i == 0 {
			if htmlString, ok := item.Html(); ok == nil {
				match := regular.ExtractAllString([]byte(htmlString), regular.TodayRe)
				result.Today = string(match[0][1])
				result.Date = strings.Replace(string(match[0][2]), " ", "", -1)
				result.Week = strings.Replace(string(match[0][3]), " ", "", -1)
			}
		} else if i == 1 {
			if htmlString, ok := item.Html(); ok == nil {
				match := regular.ExtractString([]byte(htmlString), regular.DescImageRe)
				result.AdviceImg = match
			}
		} else if i == 2 {
			if htmlString, ok := item.Html(); ok == nil {
				match := regular.ExtractAllString([]byte(htmlString), regular.CommonRe)
				for index, value := range match {
					if len(string(value[1])) > 1 {
						if index == 0 {
							result.Situation = string(value[1])
						} else if index == 2 {
							result.Temperature = string(value[1])
						} else if index == 5 {
							result.Wind = string(value[1])
						} else if index == 6 {
							result.Rainfall = string(value[1])
						}
					}
				}
			}
		} else if i == 4 {
			for i, v := range strings.Split(item.Text(), "\n") {
				if i == 0 {
					result.BodyTemperature = v
				}
			}

			if htmlString, ok := item.Html(); ok == nil {
				match := regular.ExtractAllString([]byte(htmlString), regular.CommonRe)
				for index, value := range match {
					if len(string(value[1])) > 1 {
						if index == 0 {
							result.Wet = string(value[1])
						} else if index == 1 {
							result.Press = string(value[1])
						} else if index == 2 {
							result.Visibility = string(value[1])
						}
					}
				}
			}
		}
	})

	result.SubTitle = dom.Find("div.maptop1").Eq(0).Text()

	html, err := dom.Find("div.sm").Html()
	if err != nil {
		return nil, err
	}

	update := regular.ExtractString([]byte(html), regular.TimeRe)
	result.Update = update

	return &result, nil
}

func NearDay3(content []byte) (*string, error) {
	var result string
	dom, err := regular.NewDom(content)

	if err != nil {
		return nil, err
	}

	dom.Find("div.rbox_c > div.tianqi").Each(func(i int, item *goquery.Selection) {
		if html, err := item.Html(); err == nil {
			imgDom, err := regular.NewDom([]byte(html))
			if err == nil {
				imgDom.Find("img").Each(func(i int, img *goquery.Selection) {
					if image, ok := img.Attr("src"); ok {
						html = strings.Replace(html, `src="`+image, `src="`+"https://www.qixiangwang.cn"+image, 1)
					}
				})
			}
			result = html
		}
	})

	return &result, err
}

func Exponent(content []byte) (*string, error) {
	var result string
	dom, err := regular.NewDom(content)
	if err != nil {
		return nil, err
	}

	dom.Find("div.bbox_c3").Each(func(i int, item *goquery.Selection) {
		if html, err := item.Html(); err == nil {
			result = html
		}
	})
	return &result, nil
}

func Pm25(content []byte) (*models.Pm25, error) {
	var result models.Pm25
	dom, err := regular.NewDom(content)
	if err != nil {
		return nil, err
	}

	// 城市名
	dom.Find("div.city_name > h2").Each(func(i int, item *goquery.Selection) {
		result.CityName = item.Text()
	})

	dom.Find("div.level > h4").Each(func(i int, item *goquery.Selection) {
		result.AirQuality = strings.Replace(item.Text(), " ", "", -1)
	})

	dom.Find("div.affect").Each(func(i int, item *goquery.Selection) {
		tex := strings.Replace(item.Text(), "\n", "", -1)
		result.Health = strings.Replace(tex, " ", "", -1)
	})

	dom.Find("div.action").Each(func(i int, item *goquery.Selection) {
		tex := strings.Replace(item.Text(), "\n", "", -1)
		result.Advice = strings.Replace(tex, " ", "", -1)
	})

	return &result, nil
}