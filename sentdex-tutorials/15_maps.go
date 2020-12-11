package main

/*import ("fmt"
"net/http"
"io/ioutil"
"encoding/xml")*/

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

var requestedXML = []byte(`
<sitemapindex>
	<sitemap>
		<loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
	</sitemap>

</sitemapindex>`)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

func main() {
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

	for idx, data := range newsMap {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}
}
