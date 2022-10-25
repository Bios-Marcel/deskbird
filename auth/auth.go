package auth

import (
	"context"
	"os"
	"path/filepath"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func RetrieveBearerToken() (string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	authReceiver := make(chan string, 1)

	// Default flags from chromedp.DefaultExecAllocatorOptions with modifications.
	options := []chromedp.ExecAllocatorOption{
		// Required in order to prevent having to log in over and over again.
		chromedp.UserDataDir(filepath.Join(userConfigDir, "deskbird_robot")),

		// Defaults:

		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		//Headless has been removed here

		// After Puppeteer's default behavior.
		chromedp.Flag("disable-background-networking", true),
		chromedp.Flag("enable-features", "NetworkService,NetworkServiceInProcess"),
		chromedp.Flag("disable-background-timer-throttling", true),
		chromedp.Flag("disable-backgrounding-occluded-windows", true),
		chromedp.Flag("disable-breakpad", true),
		chromedp.Flag("disable-client-side-phishing-detection", true),
		chromedp.Flag("disable-default-apps", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("disable-features", "site-per-process,Translate,BlinkGenPropertyTrees"),
		chromedp.Flag("disable-hang-monitor", true),
		chromedp.Flag("disable-ipc-flooding-protection", true),
		chromedp.Flag("disable-popup-blocking", true),
		chromedp.Flag("disable-prompt-on-repost", true),
		chromedp.Flag("disable-renderer-backgrounding", true),
		chromedp.Flag("disable-sync", true),
		chromedp.Flag("force-color-profile", "srgb"),
		chromedp.Flag("metrics-recording-only", true),
		chromedp.Flag("safebrowsing-disable-auto-update", true),
		chromedp.Flag("enable-automation", true),
		chromedp.Flag("password-store", "basic"),
		chromedp.Flag("use-mock-keychain", true),
	}

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch specificEvent := ev.(type) {
		case *network.EventRequestWillBeSent:
			auth := specificEvent.Request.Headers["Authorization"]
			if auth != nil {
				chromedp.Cancel(ctx)
				authReceiver <- auth.(string)
			}
		}
	})

	err = chromedp.Run(ctx, chromedp.Navigate("https://app.deskbird.com/"))
	if err != nil {
		return "", err
	}

	bearerToken := <-authReceiver
	close(authReceiver)
	return bearerToken, nil
}
