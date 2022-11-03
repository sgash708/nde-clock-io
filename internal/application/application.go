package application

import (
	"log"

	"github.com/sgash708/nde-clock-io/internal/client/chromedriver"
	"github.com/sgash708/nde-clock-io/internal/client/slack"
	"github.com/sgash708/nde-clock-io/internal/config"
	"github.com/sgash708/nde-clock-io/internal/util"
	"golang.org/x/xerrors"
)

const NoTimeMsg = "not exist time."

func RunClockIn() (err error) {
	d := chromedriver.StartDriver(config.Conf.Secret.URL)

	defer func() {
		var err error
		if err = d.Stop(); err != nil {
			return
		}
	}()

	site := diChromeDriver()
	if err := site.Login(); err != nil {
		return err
	}
	now := site.GetTime()
	if now == "" {
		return xerrors.New(NoTimeMsg)
	}
	if err := site.ClockIn(); err != nil {
		return err
	}

	fileName := util.GetTimeFileName()
	if err := site.ScreenShot(fileName); err != nil {
		return err
	}
	sl := diSlack()
	if err := sl.UploadFile(fileName); err != nil {
		return err
	}
	log.Printf("clock in: %v\n", now)

	return nil
}

func RunClockOut() (err error) {
	d := chromedriver.StartDriver(config.Conf.Secret.URL)

	defer func() {
		var err error
		if err = d.Stop(); err != nil {
			return
		}
	}()

	site := diChromeDriver()
	if err := site.Login(); err != nil {
		return err
	}
	now := site.GetTime()
	if now == "" {
		return xerrors.New(NoTimeMsg)
	}
	if err := site.ClockOut(); err != nil {
		return err
	}

	fileName := util.GetTimeFileName()
	if err := site.ScreenShot(fileName); err != nil {
		return err
	}
	sl := diSlack()
	if err := sl.UploadFile(fileName); err != nil {
		return err
	}
	log.Printf("clock out: %v\n", now)

	return nil
}

func diChromeDriver() chromedriver.ISite {
	return chromedriver.NewSite(
		config.Conf.Secret.Password,
		config.Conf.Secret.UserID,
	)
}

func diSlack() slack.SlackInterface {
	return slack.NewSlack(
		config.Conf.Secret.Token,
		config.Conf.Secret.Channel,
	)
}
