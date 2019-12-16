package chromedriver

type ErrStart struct{ cause error }

func (e ErrStart) Error() string {
	return "error starting chromedriver: " + e.cause.Error()
}

func (e ErrStart) Unwrap() error {
	return e.cause
}

type ErrStop struct{ cause error }

func (e ErrStop) Error() string {
	return "error stopping chromedriver: " + e.cause.Error()
}

func (e ErrStop) Unwrap() error {
	return e.cause
}
