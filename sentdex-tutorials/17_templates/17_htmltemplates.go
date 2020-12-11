package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var requestedXML = []byte(`
<sitemapindex>
	<sitemap>
		<loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
	</sitemap>

</sitemapindex>`)

type NewsMap struct {
	Keyword  string
	Location string
}

type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {

	var s SitemapIndex
	var n News
	newsMap := make(map[string]NewsMap)
	//	resp, _ := http.Get("http://www.washingtonpost.com/news-sitemap-index.xml")
	//	bytes, _ := ioutil.ReadAll(resp.Body)
	//	resp.Body.Close()

	//since the code of wp changed, we won't make the request
	bytes := requestedXML
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		for idx, _ := range n.Keywords {
			newsMap[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	p := NewsAggPage{Title: "Amazing News Aggregator", News: newsMap}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Web!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":9954", nil)
}
