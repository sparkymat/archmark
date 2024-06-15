package http

import httplib "net/http"

type Client interface {
	Do(request *httplib.Request) (*httplib.Response, error)
}
