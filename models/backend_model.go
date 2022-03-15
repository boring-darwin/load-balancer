package models

import (
	"net/http/httputil"
)

type Backend struct {
	Url     string
	Proxy   *httputil.ReverseProxy
	Healthy bool
}

type BackendList struct {
	BL []*Backend
}
