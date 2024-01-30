package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

type Product struct {
	Categories []string            `json:"categories"`
	Name       string              `json:"name"`
	Price      string              `json:"price"`
	Available  string              `json:"available"`
	Sold       string              `json:"sold"`
	ImageLinks []string            `json:"image_links"`
	Specs      map[string]string   `json:"specs"`
	Attributes map[string][]string `json:"attributes"`
}

func main() {

	file, err := os.Create("clothes_women.jsonl")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Create a buffered writer
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*.ebay.com",
		Delay:       1 * time.Second,
		RandomDelay: 1 * time.Second,
	})

	c.CacheDir = "./ebay_cache"
	extensions.RandomUserAgent(c)

	detailCollector := c.Clone()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error: ", err)
	})

	c.OnHTML("body", func(h *colly.HTMLElement) {
		h.ForEach("section.b-module.b-carousel.b-display--landscape:first-of-type", func(i int, g *colly.HTMLElement) {

			title := g.ChildText("h2.section-title__title")
			if title == "" {
				return
			}

			fmt.Println(title)
			if title == "Shop by Category" {
				g.ForEach("a.b-guidancecard__link", func(i int, d *colly.HTMLElement) {
					link := d.Attr("href")
					c.Visit(link)
				})
			} else {
				h.ForEach("a.s-item__link", func(i int, f *colly.HTMLElement) {
					if i < 10 {
						link := f.Attr("href")
						link = strings.Split(link, "?hash")[0]
						println(link)
						detailCollector.Visit(link)
					}
				})
			}
		})
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited: ", r.Request.URL)
	})

	// this is single product page
	detailCollector.OnHTML(".vim.x-vi-evo-main-container.template-evo-avip", func(e *colly.HTMLElement) {
		fmt.Println("Started scraping product page")

		product := Product{}

		// Step 1. Scrape category breadcrumb
		var categories []string
		e.ForEach("nav.breadcrumbs li > a > span", func(_ int, h *colly.HTMLElement) {
			categories = append(categories, h.Text)
		})

		if len(categories) == 0 {
			fmt.Println("Failed to scrape product categories")
			return
		}

		product.Categories = categories
		// Step 2. Scrape product name
		productName := e.ChildText("h1.x-item-title__mainTitle > span.ux-textspans.ux-textspans--BOLD")
		if productName == "" {
			fmt.Println("Failed to scrape product name")
			return
		}

		product.Name = productName

		// Step 3. Scrape product prices
		price := e.ChildText("div.x-price-primary > span")

		if price == "" {
			fmt.Println("Failed to scrape product price")
			return
		}

		product.Price = price

		// Step 4. Scrape available and sold counts
		available := e.ChildText("div.d-quantity__availability > div > span:first-child")
		sold := e.ChildText("div.d-quantity__availability > div > span:last-child")

		product.Available = available
		product.Sold = sold

		// Step 5. Scrape product variants if exists
		productAttributs := make(map[string][]string)
		e.ForEach("label.x-msku__label", func(i int, h *colly.HTMLElement) {
			attribute := h.ChildText("span.x-msku__label-text > span")
			var values []string
			h.ForEach("span.x-msku__select-box-wrapper > select > option:not(:first-child)", func(i int, j *colly.HTMLElement) {
				values = append(values, j.Text)
			})
			productAttributs[attribute] = values
		})

		product.Attributes = productAttributs

		// Step 6. Scrape image links
		var imgLinks []string
		e.ForEach("div.ux-image-carousel-container div[tabindex='0'] div.ux-image-carousel-item.image-treatment.image > img", func(i int, h *colly.HTMLElement) {
			ds := h.Attr("data-src")
			if ds != "" {
				imgLinks = append(imgLinks, ds)
				return
			}
			s := h.Attr("src")
			if s != "" {
				imgLinks = append(imgLinks, s)
				return
			}

		})
		product.ImageLinks = imgLinks

		// Step 7. Scrape product specs
		productSpecs := make(map[string]string)
		e.ForEach("div.vim.x-about-this-item div.ux-layout-section-evo__col", func(i int, h *colly.HTMLElement) {
			label := h.ChildText("div.ux-labels-values__labels-content span.ux-textspans")
			value := h.ChildText("div.ux-labels-values__values-content span.ux-textspans")
			productSpecs[label] = value
		})

		product.Specs = productSpecs

		jsonData, err := json.Marshal(product)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}

		// Write JSON line to file
		_, err = writer.WriteString(string(jsonData) + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)

	})

	c.Visit("https://www.ebay.com/b/TV-Video-Home-Audio-Electronics/32852/bn_1648392")

}
