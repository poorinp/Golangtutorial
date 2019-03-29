package main

import ("fmt"
// "net/http"
// "io/ioutil"
"encoding/xml")

var washPostXML = []byte(`
<sitemapindex>
	<sitemap>
		<loc>https://www.washingtonpost.com/news-sitemaps/politics.xml</loc>		
	</sitemap>
	<sitemap>
		<loc>https://www.washingtonpost.com/news-sitemaps/opinions.xml</loc>
	</sitemap>
	<sitemap>
		<loc>https://www.washingtonpost.com/news-sitemaps/business.xml</loc>		
	</sitemap>
</sitemapindex>`)

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"'`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	// first section when use ioutil
	// resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	// bytes, _ := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()

	// var s SitemapIndex
	// xml.Unmarshal(bytes, &s)
	// fmt.Println(s.Locations)

	//second section not use ioutil
	// resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes := washPostXML

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
}