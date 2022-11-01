package chromedriver

import (
	"testing"
)

func TestStartDriver(t *testing.T) {
	cases := []struct {
		url string
	}{
		{url: "https://www.google.com"},
		{url: "https://www.yahoo.co.jp"},
	}

	for _, c := range cases {
		driver := StartDriver(c.url)
		if err := driver.Stop(); err != nil {
			t.Error("ドライバを終了させることができませんでした。起動していないことが想定されます。")
		}
	}
}
