package regular

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

var (
	TodayRe     = regexp.MustCompile(`<b>([^<]+)<\/b><br\/>([^<]+)<br\/>([^<]+)`)
	DescImageRe = regexp.MustCompile(`<img.*src="([^"]+)"[^>]*>`)
	CommonRe    = regexp.MustCompile(`>\s*([^<]+)`)
	TimeRe    = regexp.MustCompile(`发布日期：([^<]+)`)
	HotNewsRe = regexp.MustCompile(`<i>[^<]+</i>\s+<a href="([^"]+)">([^<]+)</a>`)
	//CodeRe = regexp.MustCompile(`([1-8][1-7]\d{4}(?:199\d|200[01])(?:0[1-9]|1[0-2])(?:0[1-9]|[12]\d|3[01])\d{3}[\dX])`)
	CodeRe = regexp.MustCompile(`([1-8][1-7]\d{4}(?:1991)(?:0[1-9]|1[0-2])(?:0[1-9]|[12]\d|3[01])\d{3}[\dX])`)
)

func ExtractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

func ExtractAllString(contents []byte, re *regexp.Regexp) [][][]byte {
	return re.FindAllSubmatch(contents, -1)
}

func NewDom(content []byte) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(strings.NewReader(string(content)))
}
