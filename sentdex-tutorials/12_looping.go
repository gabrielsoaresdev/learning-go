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
	<sitemap>
		<loc>http://www.washingtonpost.com/news-options-sitemap.xml</loc>
	</sitemap>
</sitemapindex>`)

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
//	resp, _ := http.Get("http://www.washingtonpost.com/news-sitemap-index.xml")
//	bytes, _ := ioutil.ReadAll(resp.Body)
//	resp.Body.Close()

	//since the code of wp changed, we won't make the request
	bytes := requestedXML

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)

}