package engine

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Shopify/go-lua"
)

type httpBind struct {
	client *http.Client
}

func newHTTPBinding() *httpBind {
	return &httpBind{
		client: &http.Client{},
	}
}

func (h *httpBind) get(l *lua.State) int {

	u := lua.CheckString(l, -1)

	resp, err := h.client.Get(u)
	if err != nil {
		lua.Errorf(l, "lua-http: can't GET: %v", err)
		return 0
	}
	defer resp.Body.Close()

	args, err := pushResponse(l, resp)
	if err != nil {
		lua.Errorf(l, "lua-http: can't read body from GET: %v", err)
		return args
	}

	return args
}

func (h *httpBind) post(l *lua.State) int {

	u := lua.CheckString(l, -3)
	contentType := lua.CheckString(l, -2)
	body := lua.CheckString(l, -1)

	resp, err := h.client.Post(u, contentType, strings.NewReader(body))
	if err != nil {
		lua.Errorf(l, "lua-http: can't POST: %v", err)
		return 0
	}
	defer resp.Body.Close()

	args, err := pushResponse(l, resp)
	if err != nil {
		lua.Errorf(l, "lua-http: can't read body from POST: %v", err)
		return args
	}

	return args
}

func pushResponse(l *lua.State, resp *http.Response) (int, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	l.NewTable()
	setstring := func(key, value string) {
		l.PushString(value)
		l.SetField(-2, key)
	}

	setfloat := func(key string, value float64) {
		l.PushNumber(value)
		l.SetField(-2, key)
	}

	// push the basic fields
	setfloat("code", float64(resp.StatusCode))
	setstring("status", string(resp.Status))
	setstring("body", string(body))
	setfloat("content_length", float64(resp.ContentLength))

	// push the header
	l.NewTable()
	for k := range resp.Header {
		setstring(k, resp.Header.Get(k))
	}
	l.SetField(-2, "header")

	return 1, nil
}
