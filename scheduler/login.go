package scheduler

import (
	"github.com/fedesog/webdriver"
	"github.com/ycombinator/courtgrabber/chromedriver"
)

func Login(username, password string) error {
	session, err := chromedriver.Session()
	if err != nil {
		return ErrLogin{err}
	}

	err = session.Url(loginURL)
	if err != nil {
		return ErrLogin{err}
	}

	err = enterUsername(session, username)
	if err != nil {
		return ErrLogin{err}
	}

	err = enterPassword(session, password)
	if err != nil {
		return ErrLogin{err}
	}

	err = submit(session)
	if err != nil {
		return ErrLogin{err}
	}

	return nil
}

func enterUsername(session *webdriver.Session, username string) error {
	usernameEl, err := session.FindElement(webdriver.ID, "j_username")
	if err != nil {
		return err
	}

	err = usernameEl.SendKeys(username)
	if err != nil {
		return err
	}

	return nil
}

func enterPassword(session *webdriver.Session, password string) error {
	passwordEl, err := session.FindElement(webdriver.ID, "j_password")
	if err != nil {
		return err
	}

	err = passwordEl.SendKeys(password)
	if err != nil {
		return err
	}

	return nil
}

func submit(session *webdriver.Session) error {
	submitEl, err := session.FindElement(webdriver.ClassName, "btn-primary")
	if err != nil {
		return err
	}

	err = submitEl.Submit()
	if err != nil {
		return err
	}

	return nil
}
