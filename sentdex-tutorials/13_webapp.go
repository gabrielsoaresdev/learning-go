package main

/*import ("fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml")*/

import ("fmt"
	"encoding/xml")

var requestedXML = []byte(`
<sitemapindex>
	<sitemap>
		<loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
	</sitemap>
	<sitemap>
		<loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
	</sitemap>
</sitemapindex>`)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>Keywords"`
	Location []string `xml:"url>loc"`
}

func main() {
	var s SitemapIndex
	var n News
//	resp, _ := http.Get("http://www.washingtonpost.com/news-sitemap-index.xml")
//	bytes, _ := ioutil.ReadAll(resp.Body)
//	resp.Body.Close()

	//since the code of wp changed, we won't make the request
	bytes := requestedXML
	xml.Unmarshal(bytes, &s)

	for i, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
	}
}