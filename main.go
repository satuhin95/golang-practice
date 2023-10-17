package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
)

type item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func main() {
	app := fiber.New()
	app.Get("/scrape", func(c *fiber.Ctx) error {
		var items []item
		collector := colly.NewCollector(
			colly.AllowedDomains("j2store.net"),
		)
		collector.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		})

		collector.OnHTML("div.col-sm-9 div[itemprop=itemListElement] ", func(h *colly.HTMLElement) {
			item := item{
				Name:   h.ChildText("h2.product-title"),
				Price:  h.ChildText("div.sale-price"),
				ImgUrl: h.ChildAttr("img", "src"),
			}
			items = append(items, item)
		})

		collector.Visit("http://j2store.net/demo/index.php/shop") // initiate a request to the specified URL.
		return c.JSON(items)                                      //we return the extracted data to the client by calling the c.JSON(...) method.
	})

	app.Listen(":8080")
}
