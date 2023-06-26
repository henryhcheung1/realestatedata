package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type PokemonProduct struct {
	url, image, name, price string
}

func main() {
	var pokemonProducts []PokemonProduct

	// initializing a chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// navigate to the target web page and select the HTML elements of interest
	var nodes []*cdp.Node
	chromedp.Run(ctx,
		chromedp.Navigate("https://scrapeme.live/shop"),
		chromedp.Nodes(".product", &nodes, chromedp.ByQueryAll),
	)

	for i := 0; i < len(nodes); i++ {
		fmt.Println(nodes[i])
	}
	fmt.Println(len(nodes))

	// scraping data from each node
	var url, image, name, price string
	for _, node := range nodes {
		chromedp.Run(ctx,
			chromedp.AttributeValue("a", "href", &url, nil, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.AttributeValue("img", "src", &image, nil, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text("h2", &name, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(".price", &price, chromedp.ByQuery, chromedp.FromNode(node)),
		)

		pokemonProduct := PokemonProduct{}

		pokemonProduct.url = url
		pokemonProduct.image = image
		pokemonProduct.name = name
		pokemonProduct.price = price

		pokemonProducts = append(pokemonProducts, pokemonProduct)
	}

	for _, pokemon := range pokemonProducts {
		fmt.Println(pokemon)
	}

	// export logic
}

// // -------------------------------------------------------
// // Example to scrape and traverse links
// // Reference: https://devmarkpro.com/chromedp-working-with-nodes-and-tabs
// package main

// import (
//     "context"
//     "fmt"
//     "log"

//     "github.com/chromedp/cdproto/cdp"
//     "github.com/chromedp/chromedp"
// )

// func main() {
//     ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithErrorf(log.Printf))
//     defer cancel()
//     var nodes []*cdp.Node
//     selector := "#main ul li a"
//     pageURL := "https://notepad-plus-plus.org/downloads/"
//     if err := chromedp.Run(ctx, chromedp.Tasks{
//         chromedp.Navigate(pageURL),
//         chromedp.WaitReady(selector),
//         chromedp.Nodes(selector, &nodes),
//     }); err != nil {
//         panic(err)
//     }
//     f := func(ctx context.Context, url string) {
//         clone, cancel := chromedp.NewContext(ctx)
//         defer cancel()
//         fmt.Printf("%s is opening in a new tab\n", url)

//         if err := chromedp.Run(clone, chromedp.Navigate(url)); err != nil {
//             // do something nice with you errors!
//             panic(err)
//         }
//     }
//     for _, n := range nodes {
//         u := n.AttributeValue("href")
//         go f(ctx, u)
//     }
// }

// -------------------------------------------------------

// Example to download source html
// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/chromedp/cdproto/emulation"
// 	"github.com/chromedp/chromedp"
// )

// func loadHTML() {
// 	// Create a new context
// 	ctx, cancel := chromedp.NewContext(context.Background())
// 	defer cancel()

// 	// Set up a timeout for the context
// 	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
// 	defer cancel()

// 	// Create a variable to store the HTML content
// 	var htmlContent string

// 	// Run the scraping task
// 	err := chromedp.Run(ctx,
// 		emulation.SetUserAgentOverride("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"),
// 		chromedp.Navigate("https://www.zillow.com/irvine-ca/?searchQueryState=%7B%22usersSearchTerm%22%3A%22Los%20Angeles%2C%20CA%22%2C%22mapBounds%22%3A%7B%22north%22%3A33.82412885475799%2C%22east%22%3A-117.67383009448243%2C%22south%22%3A33.54974119236114%2C%22west%22%3A-117.87467390551758%7D%2C%22isMapVisible%22%3Atrue%2C%22filterState%22%3A%7B%22sort%22%3A%7B%22value%22%3A%22days%22%7D%2C%22ah%22%3A%7B%22value%22%3Atrue%7D%2C%22sche%22%3A%7B%22value%22%3Afalse%7D%2C%22schm%22%3A%7B%22value%22%3Afalse%7D%2C%22schh%22%3A%7B%22value%22%3Afalse%7D%2C%22schp%22%3A%7B%22value%22%3Afalse%7D%2C%22schr%22%3A%7B%22value%22%3Afalse%7D%2C%22schc%22%3A%7B%22value%22%3Afalse%7D%2C%22schu%22%3A%7B%22value%22%3Afalse%7D%2C%22mf%22%3A%7B%22value%22%3Afalse%7D%2C%22land%22%3A%7B%22value%22%3Afalse%7D%2C%22apa%22%3A%7B%22value%22%3Afalse%7D%2C%22manu%22%3A%7B%22value%22%3Afalse%7D%7D%2C%22isListVisible%22%3Atrue%2C%22mapZoom%22%3A13%2C%22regionSelection%22%3A%5B%7B%22regionId%22%3A52650%2C%22regionType%22%3A6%7D%5D%2C%22pagination%22%3A%7B%7D%7D"),
// 		chromedp.WaitVisible("body", chromedp.ByQuery),
// 		chromedp.OuterHTML("html", &htmlContent),
// 	)
// 	if err != nil {
// 		log.Fatal("Failed to load HTML:", err)
// 	}

// 	// Print the HTML content
// 	fmt.Println(htmlContent)
// }

// func main() {
// 	loadHTML()
// }
