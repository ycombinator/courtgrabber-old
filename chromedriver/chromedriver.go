package chromedriver

import (
	"os"
	"path"

	"github.com/fedesog/webdriver"
)

const ()

var driver *webdriver.ChromeDriver
var session *webdriver.Session

func Session() (*webdriver.Session, error) {
	err := startDriver()
	if err != nil {
		return nil, ErrStart{err}
	}

	session, err = startSession()
	if err != nil {
		return nil, ErrStart{err}
	}

	return session, nil
}

func Close() error {
	err := deleteSession()
	if err != nil {
		return ErrStop{err}
	}

	err = stopDriver()
	if err != nil {
		return ErrStop{err}
	}

	return nil
}

func startDriver() error {
	if driver != nil {
		return nil
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	driver = webdriver.NewChromeDriver(path.Join(cwd, "chromedriver", "chromedriver"))
	err = driver.Start()
	if err != nil {
		return err
	}

	return nil
}

func startSession() (*webdriver.Session, error) {
	if session != nil {
		return session, nil
	}

	desired := webdriver.Capabilities{}
	required := webdriver.Capabilities{}

	var err error
	session, err = driver.NewSession(desired, required)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func stopDriver() error {
	if driver == nil {
		return nil
	}

	err := driver.Stop()
	if err != nil {
		return err
	}

	return nil
}

func deleteSession() error {
	if session == nil {
		return nil
	}

	err := session.Delete()
	if err != nil {
		return err
	}

	return nil
}
