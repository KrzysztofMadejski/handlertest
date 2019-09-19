package assert

func (a *Assert) Header(key string, value string) *Assert {
	values := a.R.Header[key]
	if values == nil || len(values) == 0 {
		a.T.Errorf("Expected header %s to be set, it is not", key)

	} else if got := a.R.Header.Get(key); got != value {
		a.T.Errorf("Expected header %s to be set to '%s', got '%s'", key, value, got)
	}

	return a
}

func (a *Assert) HeaderMissing(key string) *Assert {
	value := a.R.Header.Get(key)
	if value != "" {
		a.T.Errorf("Expected header %s to be empty, got '%s'", key, value)
	}

	return a
}

func (a *Assert) ContentType(contentType string) *Assert {
	return a.Header("Content-Type", contentType)
}
