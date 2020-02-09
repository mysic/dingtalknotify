package ding

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	httpMethodPost = "POST"
	httpMethodGet  = "GET"
)

type requestHandler struct {
	scheme      string
	host        string
	path        string
	charset     string
	contentType string
	method      string
	queryString map[string]interface{}
}

func (h *requestHandler) post(url string, data *map[string]interface{}) (string, error) {
	body, _ := MapToJson(*data)
	return h._send(url, httpMethodPost, body)
}

func (h *requestHandler) get(url string) (string, error) {
	return h._send(url, httpMethodGet, &[]byte{})
}

func (h *requestHandler) _send(url string, method string, payload *[]byte) (r string, err error) {
	for {
		if url == "" {
			err = errors.New("url can't be empty")
			break
		}
		if method == "" {
			err = errors.New("method can't be empty")
			break
		}
		if method == httpMethodPost && len(*payload) == 0 {
			err = errors.New("payload can't be empty")
			break
		}

		req, err := http.NewRequest(method, url, bytes.NewBuffer(*payload))
		if err != nil {
			break
		}
		req.Header.Add("Content-Type", h.contentType)
		req.Header.Add("Charset", h.charset)
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			if resp != nil {
				err = resp.Body.Close()
			}
			break
		}
		if resp == nil {
			err = errors.New("response body is empty")
			break
		}
		defer resp.Body.Close()
		res, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			break
		}
		r = string(res)
		break
	}

	return r, err
}

func (h *requestHandler) makeUrl() (theUrl string, err error) {
	for {
		if h.path == "" {
			err = errors.New("path cant't be empty")
			break
		}
		if len(h.queryString) == 0 {
			err = errors.New("queryString can't be empty")
			break
		}
		v := &url.Values{}
		for key, value := range h.queryString {
			v.Add(key, ToString(value))
		}
		query := v.Encode()
		u := &url.URL{
			Scheme:   h.scheme,
			Host:     h.host,
			Path:     h.path,
			RawQuery: query,
		}
		theUrl = u.String()

		break
	}
	return theUrl, err
}

func (h *requestHandler) setQueryString(queryString *map[string]interface{}) *requestHandler {
	for key, val := range *queryString {
		h.queryString[key] = val
	}
	return h
}
