package chromedriver

import (
	"fmt"
	"strings"
	"time"
)

type ISite interface {
	Login() error
	IsHoliday() (bool, error)
	ClockIn() error
	ClockOut() error
	GetTime() string
	ScreenShot(string) error
	DoBackTopPage() error
}

type Site struct {
	Password string
	UserID   string
}

func NewSite(password, userID string) ISite {
	return &Site{
		Password: password,
		UserID:   userID,
	}
}

func (s *Site) Login() error {
	time.Sleep(3 * time.Second)
	if err := page.FindByXPath(`//*[@id="txtID"]`).Fill(s.UserID); err != nil {
		return err
	}
	if err := page.FindByXPath(`//*[@id="txtPsw"]`).Fill(s.Password); err != nil {
		return err
	}
	if err := page.FindByXPath(`//*[@id="btnLogin"]`).Click(); err != nil {
		return err
	}
	time.Sleep(2 * time.Second)

	if err := page.FindByXPath(`//*[@id="ctl00_ContentPlaceHolder1_imgBtnSyuugyou"]`).Click(); err != nil {
		return err
	}
	time.Sleep(2 * time.Second)

	return nil
}

func (s *Site) IsHoliday() (bool, error) {
	isHoliday := false
	if err := page.FindByXPath(`//div[4]/div[2]/a`).Click(); err != nil {
		return isHoliday, err
	}
	time.Sleep(3 * time.Second)

	day := time.Now().Day()
	dayTr := fmt.Sprintf(`table:nth-child(1) > tbody > tr:nth-child(%d)`, day+1)
	content := GetHTML()
	// company holiday
	colorCode, isExist := content.Find(dayTr + `> td:nth-child(2)`).Attr(`bgcolor`)
	if isExist && colorCode == `#FF00FF` {
		isHoliday = true
		return isHoliday, nil
	}
	// my holiday
	remarksText := content.Find(dayTr + `> td:nth-child(7)`).Text()
	if strings.Contains(remarksText, "有休") {
		isHoliday = true
		return isHoliday, nil
	}
	fmt.Println(colorCode, remarksText)

	return isHoliday, nil
}

func (s *Site) ClockIn() error {
	if err := page.FindByXPath(`//*[@id="ctl00_ContentPlaceHolder1_ibtnIn3"]`).Click(); err != nil {
		return err
	}
	time.Sleep(5 * time.Second)

	return nil
}

func (s *Site) ClockOut() error {
	if err := page.FindByXPath(`//*[@id="ctl00_ContentPlaceHolder1_ibtnOut4"]`).Click(); err != nil {
		return err
	}
	time.Sleep(5 * time.Second)

	return nil
}

func (s *Site) GetTime() string {
	content := GetHTML()
	return content.Find(`#ctl00_ContentPlaceHolder1_lblHour`).Text()
}

func (s *Site) ScreenShot(fileName string) error {
	return page.Screenshot(fileName)
}

func (s *Site) DoBackTopPage() error {
	if err := page.FindByXPath(`//div/div[1]/div[3]/div/a`).Click(); err != nil {
		return err
	}
	time.Sleep(2 * time.Second)

	return nil
}
