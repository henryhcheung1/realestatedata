package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
)

type Property struct {
	url, price, address, size, lotSize, bedCount, bathCount string
}

// #main > ul > li.post-759.product.type-product.status-publish.has-post-thumbnail.product_cat-pokemon.product_cat-seed.product_tag-bulbasaur.product_tag-overgrow.product_tag-seed.first.instock.sold-individually.taxable.shipping-taxable.purchasable.product-type-simple
// #grid-search-results > ul > li:nth-child(42)
//*[@id="grid-search-results"]/ul/li[4]
// #zpid_2057006194 > div > div.StyledPropertyCardDataWrapper-c11n-8-89-0__sc-1omp4c3-0.eBcgiS.property-card-data > div.StyledPropertyCardDataArea-c11n-8-89-0__sc-yipmu-0.eLdkcJ > ul > li:nth-child(1) >
// #zpid_2057006194 > div > div.StyledPropertyCardDataWrapper-c11n-8-89-0__sc-1omp4c3-0.eBcgiS.property-card-data > div.StyledPropertyCardDataArea-c11n-8-89-0__sc-yipmu-0.eLdkcJ > ul > li:nth-child(2) > b
// #zpid_2057006194 > div > div.StyledPropertyCardDataWrapper-c11n-8-89-0__sc-1omp4c3-0.eBcgiS.property-card-data > div.StyledPropertyCardDataArea-c11n-8-89-0__sc-yipmu-0.eLdkcJ > ul > li:nth-child(3) > b

func main() {
	// var properties []Property

	scrapeUrl := "https://www.zillow.com/irvine-ca/?searchQueryState=%7B%22mapBounds%22%3A%7B%22north%22%3A33.81050873395913%2C%22east%22%3A-117.60636717333985%2C%22south%22%3A33.56340269357558%2C%22west%22%3A-117.94213682666016%7D%2C%22isMapVisible%22%3Atrue%2C%22filterState%22%3A%7B%22sort%22%3A%7B%22value%22%3A%22days%22%7D%2C%22ah%22%3A%7B%22value%22%3Atrue%7D%2C%22mf%22%3A%7B%22value%22%3Afalse%7D%2C%22manu%22%3A%7B%22value%22%3Afalse%7D%2C%22land%22%3A%7B%22value%22%3Afalse%7D%2C%22apa%22%3A%7B%22value%22%3Afalse%7D%7D%2C%22isListVisible%22%3Atrue%2C%22mapZoom%22%3A12%2C%22regionSelection%22%3A%5B%7B%22regionId%22%3A52650%2C%22regionType%22%3A6%7D%5D%2C%22pagination%22%3A%7B%7D%7D"

	// initializing a chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	fmt.Println("A")

	// navigate to the target web page and select the HTML elements of interest
	var nodes []*cdp.Node
	chromedp.Run(ctx,
		emulation.SetUserAgentOverride("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4472.106 Safari/537.36"),
		chromedp.Navigate(scrapeUrl),
		chromedp.Nodes("#grid-search-results > ul > li", &nodes, chromedp.ByQueryAll),
	)
	fmt.Println("C")

	// for i := 0; i < len(nodes); i++ {
	// 	fmt.Println(nodes[i])
	// }
	fmt.Println(len(nodes))

	// // scraping data from each node
	var url, price, bedCount, bathCount, size, address string
	for _, node := range nodes {

		fmt.Println(node)

		// NOTE: program hangs when chromedp cant find the elment to scrape. Currently hangs on price, bed bath count, etc.

		chromedp.Run(ctx,
			chromedp.AttributeValue("a", "href", &url, nil, chromedp.ByQuery, chromedp.FromNode(node)),
			// chromedp.Text("property-card-price", &price, chromedp.ByQuery, chromedp.FromNode(node)),
			// chromedp.Text("li:nth-child(1)", &bedCount, chromedp.ByQuery, chromedp.FromNode(node)),
			// chromedp.Text("li:nth-child(2)", &bathCount, chromedp.ByQuery, chromedp.FromNode(node)),
			// chromedp.Text("li:nth-child(3)", &size, chromedp.ByQuery, chromedp.FromNode(node)),
			// chromedp.Text("property-card-addr", &address, chromedp.ByQuery, chromedp.FromNode(node)),

			// chromedp.AttributeValue("img", "src", &image, nil, chromedp.ByQuery, chromedp.FromNode(node)),
			// chromedp.Text("h2", &name, chromedp.ByQuery, chromedp.FromNode(node)),
		)

		fmt.Println(url)
		// fmt.Println(price)
		// fmt.Println(bedCount)
		// fmt.Println(bathCount)
		// fmt.Println(size)
		// fmt.Println(address)
		fmt.Println("----")

		// pokemonProduct := PokemonProduct{}

		// pokemonProduct.url = url
		// pokemonProduct.image = image
		// pokemonProduct.name = name
		// pokemonProduct.price = price

		// pokemonProducts = append(pokemonProducts, pokemonProduct)
	}

	// for _, pokemon := range pokemonProducts {
	// 	fmt.Println(pokemon)
	// }

	// export logic
}
