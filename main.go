package main

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
)

// Only enqueue the root and paths beginning with an "a"
var rxIdx = regexp.MustCompile(`http://(www\.)?vibbo\.com/motos-de-segunda-mano-toda-espana.*o=[0-9]+.*$`)
var rxPage = regexp.MustCompile(`http://(www\.)?vibbo\.com.*/a[0-9]+.*$`)
var i = 1

//var rxOk = regexp.MustCompile(`http://duckduckgo\.com(/a.*)?$`)

type Ext struct {
	*gocrawl.DefaultExtender
}

func (e *Ext) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	//if rxIdx.MatchString(ctx.SourceURL().String()) {
	fmt.Printf("%d Visit: %s Source: %s\n", i, ctx.URL(), ctx.SourceURL())
	i = i + 1
	//}
	return nil, true
}

func (e *Ext) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	//fmt.Printf("Filtering: %s\n", ctx.NormalizedURL().String())
	url := ctx.NormalizedURL().String()
	return !isVisited && (rxPage.MatchString(url) && rxIdx.MatchString(ctx.SourceURL().String()) || rxIdx.MatchString(url))
}

func main() {
	ext := &Ext{&gocrawl.DefaultExtender{}}
	// Set custom options
	opts := gocrawl.NewOptions(ext)
	opts.CrawlDelay = 0 * time.Second
	opts.LogFlags = gocrawl.LogError
	opts.SameHostOnly = false
	opts.MaxVisits = 1000000

	c := gocrawl.NewCrawlerWithOptions(opts)
	//c.Run("http://0value.com")
	c.Run("http://www.vibbo.com/motos-de-segunda-mano-toda-espana/?ca=0_s&x=1&w=1&c=6&o=1")
}
