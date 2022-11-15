package notifications

import (
	"net/url"
	"os"
)

var fakeURL string = "https://www.url.home/path"

func secret(s string) *string {
	if len(os.Getenv("dynatrace.secrets")) > 0 {
		return &s
	}
	return nil
}

func psecret(s *string) *string {
	if len(os.Getenv("dynatrace.secrets")) > 0 {
		return s
	}
	return nil
}

func usecret(s string) *string {
	if len(os.Getenv("dynatrace.secrets")) > 0 {
		if _, e := url.ParseRequestURI(s); e != nil {
			return &fakeURL
		}
		return &s
	}
	return nil
}
