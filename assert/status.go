package assert

func (a *Assert) Status(statusCode int) *Assert {
	if a.R.StatusCode != statusCode {
		a.T.Errorf("Expected statusCode %d, got %d", statusCode, a.R.StatusCode)
	}

	return a
}
