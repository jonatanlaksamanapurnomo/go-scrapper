package product

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

// GetTokopediaProducts scrapes products from the given Tokopedia URL.
func (rsc Resource) GetTokopediaProducts(ctx context.Context, url string) (resp []Product, err error) {
	allocatorCtx, cancel := chromedp.NewExecAllocator(ctx,
		append(chromedp.DefaultExecAllocatorOptions[:],
			// Run in headless mode
			chromedp.Flag("headless", true),
			// Set custom window size
			chromedp.Flag("window-size", "1920,1080"),
			// Set custom user-agent
			chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36"),
			// Ignore HTTPS errors - use with caution
			chromedp.Flag("ignore-certificate-errors", true),
			// Disable HTTP/2 - workaround
			chromedp.Flag("ssl-version-max", "tls1.2"),
			// Disable web security - use with caution
			chromedp.Flag("disable-web-security", true),
		)...,
	)
	defer cancel()

	browserCtx, cancel := chromedp.NewContext(allocatorCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	ctxWithTimeout, cancel := context.WithTimeout(browserCtx, 30*time.Second)
	defer cancel()

	var htmlContent string
	err = chromedp.Run(ctxWithTimeout,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.OuterHTML(`html`, &htmlContent, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatalf("Failed to execute run context: %v", err)
		return resp, err
	}

	products, err := buildProductFromHTML(htmlContent)
	if err != nil {
		return resp, err
	}
	return products, nil
}
