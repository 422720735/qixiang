package router

import (
	"fmt"
	"html/template"
	"net/http"
	"qixiang/fetcher"
	"qixiang/parse/models"
	"qixiang/parse/weather"
)

const httpPort = 80

func Assemble(path string) (*models.Weather, error) {
	if path == "" {
		path = "chengdu"
	}

	response, err := fetcher.Fetcher("https://www.qixiangwang.cn/" + path + ".htm")
	if err != nil {
		return nil, err
	}

	// 今天天气实况
	today, err := weather.Today(response)
	if err != nil {
		return nil, err
	}

	// 近3天
	days3, err := weather.NearDay3(response)
	if err != nil {
		return nil, err
	}
	// 生活指数
	exponent, err := weather.Exponent(response)
	if err != nil {
		return nil, err
	}

	response, err = fetcher.Fetcher("https://www.qixiangwang.cn/pm25/" + path + "/")
	if err != nil {
		fmt.Printf("fetcher err %v", err)
		return nil, err
	}

	pm25, err := weather.Pm25(response)
	if err != nil {
		return nil, err
	}

	response, err = fetcher.Fetcher("http://www.ichong123.com/")
	if err != nil {
		return nil, err
	}

	//hot, err := pet.HotKnowledge(response)
	//if err != nil {
	//	return nil, err
	//}

	today.Days3 = *days3
	today.Exponent = *exponent
	result := models.Weather{
		CityLive: today,
		Pm25:     pm25,
		//Knowledge: hot,
	}
	return &result, nil
}

func Unescaped(x string) interface{} { return template.HTML(x) }

func HandleHtml(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("index.html")
	tmpl = tmpl.Funcs(template.FuncMap{"unescaped": Unescaped})
	// 解析指定文件生成模板对象
	tmpl, err := tmpl.ParseFiles("index.html")
	if err != nil {
		fmt.Println("create gulp failed, err:", err)
		return
	}

	res, _ := Assemble(r.URL.Query().Get("city"))
	// 利用给定数据渲染模板，并将结果写入w
	tmpl.Execute(w, &res)
}

func Star() {
	http.HandleFunc("/", HandleHtml)

	fmt.Printf("Your application is running here: http://localhost:%d\n", httpPort)
	http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil)
}
