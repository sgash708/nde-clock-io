package chromedriver

import (
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"
)

var page *agouti.Page
var driver *agouti.WebDriver

func StartDriver(url string) *agouti.WebDriver {
	driver = agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--disable-gpu",
			"--headless",
			"--no-sandbox",
			"--disable-dev-shm-usage",
			"lang=ja",
			"--disable-desktop-notifications",
			"--ignore-certificate-errors",
			"--disable-extensions",
		}),
		agouti.Debug,
	)
	err := driver.Start()
	if err != nil {
		log.Fatalf("ドライバ起動失敗\n詳細: %v", err)
	}

	page, err = driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("Chromeでページが開けません\n詳細: %v", err)
	}

	// 5000ms
	if err := page.Session().SetImplicitWait(5000); err != nil {
		log.Fatalln(err)
	}

	err = page.Navigate(url)
	if err != nil {
		log.Fatalf("該当ページに遷移できません\n詳細: %v", err)
	}

	return driver
}

func GetHTML() *goquery.Document {
	rawHTML, err := page.HTML()
	time.Sleep(300 * time.Millisecond)

	if err != nil {
		if page.Refresh() != nil {
			log.Fatalf("HTMLの再取得失敗しました\n詳細: %v", err)
		}
		time.Sleep(20 * time.Second)
		rawHTML, _ = page.HTML()
	}

	strHTML := strings.NewReader(rawHTML)
	dom, err := goquery.NewDocumentFromReader(strHTML)
	if err != nil {
		log.Fatalf("HTMLの変換に失敗しました\n詳細: %v", err)
	}

	return dom
}
