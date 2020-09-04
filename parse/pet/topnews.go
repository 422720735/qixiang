package pet

import (
	"errors"
	"fmt"
	"qixiang/fetcher"
	"qixiang/parse/models"
	"qixiang/parse/regular"
)

func HotKnowledge(content []byte) (*models.Knowledge, error) {
	dom, err := regular.NewDom(content)
	if err != nil {
		fmt.Printf("goquery err %v", err)
		return nil, err
	}

	hotHref, err := dom.Find("div.ss-news").Html()
	if err != nil {
		fmt.Printf("find html dom  failed %v", err)
		return nil, err
	}

	href := regular.ExtractString([]byte(hotHref), regular.HotNewsRe)

	// 读取移动端的详情页
	// href = strings.Replace(href, "//www", "//m", -1)
	if href == "" {
		return nil, errors.New("url is Incorrect format")
	}

	// 去读取头条信息
	resp, err := fetcher.Fetcher(href)
	if err != nil {
		return nil, err
	}

	knowledge, err := topNewsInfo(resp)
	if err == nil {
		// 获取搞笑萌图
		imageList, _ := dom.Find("div.ss-video > ul.image-list-tag").Html()
		knowledge.ImageList = imageList
	}

	return knowledge, err
}

func topNewsInfo(resp []byte) (*models.Knowledge, error) {
	dom, err := regular.NewDom(resp)
	if err != nil {
		fmt.Printf("find html dom  failed %v", err)
		return nil, err
	}

	title := dom.Find("div.article-content > h1.ac-title").Text()
	content, _ := dom.Find("div.article-content > div.ac-content").Eq(0).Html()

	result := models.Knowledge{
		Title:   title,
		Content: content,
	}

	return &result, nil
}
