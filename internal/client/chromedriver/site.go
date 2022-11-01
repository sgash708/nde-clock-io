package chromedriver

import (
	"time"
)

type ISite interface {
	Login() error
	ClockIn() error
	ClockOut() error
	GetTime() string
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
	// Fill
	if err := page.FindByXPath(`//*[@id="txtID"]`).Fill(s.UserID); err != nil {
		return err
	}
	if err := page.FindByXPath(`//*[@id="txtPsw"]`).Fill(s.Password); err != nil {
		return err
	}
	// Click
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

func (s *Site) ClockIn() error {
	if err := page.FindByXPath(`//*[@id="ctl00_ContentPlaceHolder1_ibtnIn3"]`).Click(); err != nil {
		return err
	}

	return nil
}

func (s *Site) ClockOut() error {
	//*[@id="ctl00_ContentPlaceHolder1_ibtnOut4"]
	if err := page.FindByXPath(`//*[@id="ctl00_ContentPlaceHolder1_ibtnOut4"]`).Click(); err != nil {
		return err
	}

	return nil
}

func (s *Site) GetTime() string {
	content := GetHTML()
	return content.Find(`#ctl00_ContentPlaceHolder1_lblHour`).Text()
}
