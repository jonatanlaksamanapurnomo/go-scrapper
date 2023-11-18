package product

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"net/url"
	"strconv"
	"time"
)

func (rsc Resource) GetTokopediaProducts(ctx context.Context, params TokopediaSearchParams) ([]Product, error) {
	// Construct the URL from the search parameters
	baseURL := "https://www.tokopedia.com/search"
	queryValues := url.Values{}
	queryValues.Set("q", params.Query)
	queryValues.Set("page", strconv.Itoa(int(params.Page)))
	queryValues.Set("ob", params.SortOrder)
	// Add other query parameters as needed

	searchURL := fmt.Sprintf("%s?%s", baseURL, queryValues.Encode())

	allocatorCtx, cancel := chromedp.NewExecAllocator(ctx,
		append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", true),
			chromedp.Flag("window-size", "1920,1080"),
			chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36"),
			chromedp.Flag("ignore-certificate-errors", true),
			chromedp.Flag("ssl-version-max", "tls1.2"),
			chromedp.Flag("disable-web-security", true),
		)...,
	)
	defer cancel()

	browserCtx, cancel := chromedp.NewContext(allocatorCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	ctxWithTimeout, cancel := context.WithTimeout(browserCtx, 30*time.Second)
	defer cancel()

	var products []Product

	// Navigate and wait for the initial content to load
	err := chromedp.Run(ctxWithTimeout,
		chromedp.Navigate(searchURL),
		chromedp.WaitReady(`div[data-testid='divSRPContentProducts']`),
	)
	if err != nil {
		log.Fatalf("Failed to navigate and load initial content: %v", err)
		return nil, err
	}

	// Scrolling and dynamic content loading logic
	var prevProductsCount int
	for {
		var htmlContent string
		err = chromedp.Run(ctxWithTimeout,
			chromedp.Evaluate(`window.scrollTo(0, document.body.scrollHeight);`, nil),
			chromedp.Sleep(2*time.Second), // Adjust this based on site response
			chromedp.OuterHTML(`html`, &htmlContent, chromedp.ByQuery),
		)
		if err != nil {
			log.Printf("Error during scrolling and content loading: %v", err)
			break
		}

		newProducts, err := buildProductFromHTML(htmlContent)
		if err != nil {
			log.Printf("Error building products from HTML: %v", err)
			break
		}

		// Check if no new products were loaded
		if len(newProducts) == prevProductsCount {
			break
		}
		prevProductsCount = len(newProducts)

		products = append(products, newProducts...)
	}

	return products, nil
}

func (rsc Resource) GetTokopediaProductsByURL(ctx context.Context, url string) ([]Product, error) {
	allocatorCtx, cancel := chromedp.NewExecAllocator(ctx,
		append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", true),
			chromedp.Flag("window-size", "1920,1080"),
			chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36"),
			chromedp.Flag("ignore-certificate-errors", true),
			chromedp.Flag("ssl-version-max", "tls1.2"),
			chromedp.Flag("disable-web-security", true),
		)...,
	)
	defer cancel()

	browserCtx, cancel := chromedp.NewContext(allocatorCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	ctxWithTimeout, cancel := context.WithTimeout(browserCtx, 30*time.Second)
	defer cancel()

	var products []Product

	// Navigate and wait for the initial content to load
	err := chromedp.Run(ctxWithTimeout,
		chromedp.Navigate(url),
		chromedp.WaitReady(`div[data-testid='divSRPContentProducts']`),
	)
	if err != nil {
		log.Fatalf("Failed to navigate and load initial content: %v", err)
		return nil, err
	}

	// Scrolling and dynamic content loading logic
	for {
		var htmlContent string
		err = chromedp.Run(ctxWithTimeout,
			chromedp.Evaluate(`window.scrollTo(0, document.body.scrollHeight);`, nil),
			chromedp.Sleep(2*time.Second), // Adjust this based on site response
			chromedp.OuterHTML(`html`, &htmlContent, chromedp.ByQuery),
		)
		if err != nil {
			log.Printf("Error during scrolling and content loading: %v", err)
			break
		}

		newProducts, err := buildProductFromHTML(htmlContent)
		if err != nil {
			log.Printf("Error building products from HTML: %v", err)
			break
		}

		// Check if no new products were loaded
		if len(newProducts) == 0 {
			break
		}

		products = append(products, newProducts...)
	}

	return products, nil
}
