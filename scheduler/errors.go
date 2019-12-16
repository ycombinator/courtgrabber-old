package scheduler

type ErrLogin struct{ cause error }
type ErrDay struct{ cause error }
type ErrFindCourt struct{ cause error }
type ErrCourt struct{ cause error }

func (e ErrLogin) Error() string {
	return "error logging in to AVAC scheduler: " + e.cause.Error()
}

func (e ErrLogin) Unwrap() error {
	return e.cause
}

func (e ErrDay) Error() string {
	return "error finding day on AVAC schedule: " + e.cause.Error()
}

func (e ErrDay) Unwrap() error {
	return e.cause
}

func (e ErrFindCourt) Error() string {
	return "error finding free court on AVAC schedule: " + e.cause.Error()
}

func (e ErrFindCourt) Unwrap() error {
	return e.cause
}

func (e ErrCourt) Error() string {
	return "error booking court on AVAC schedule: " + e.cause.Error()
}

func (e ErrCourt) Unwrap() error {
	return e.cause
}
