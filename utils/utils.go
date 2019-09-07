package utils

import "net/url"

func ValuesFromMap(values map[string]string) url.Values {
	v := url.Values{}
	for key, val := range values {
		v.Set(key, val)
	}
	return v
}
