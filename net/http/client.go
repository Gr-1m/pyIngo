package http

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	http.Client

	tlsverify bool // 0: Skip the SSL cert Verify
	openproxy bool // 1: Open http proxy
	proxy     string
}

var DefaultClient = &Client{}

func Get(url, proxy string, headers map[string]string, verify bool) (*Response, error) {
	// if proxy != "" set DefaultClient.openproxy = true, else set it false
	DefaultClient.openproxy = (proxy != "")
	DefaultClient.tlsverify = verify
	if !strings.Contains(proxy, "http://") {
		DefaultClient.proxy = "http://" + proxy
	} else {
		DefaultClient.proxy = proxy
	}

	return DefaultClient.Get(url, headers)
}

func (c *Client) Get(endpoint string, headers map[string]string) (resp *Response, err error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if c.openproxy {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: !c.tlsverify},
			Proxy:           func(r *http.Request) (*url.URL, error) { return url.Parse(c.proxy) },
		}
	} else {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: !c.tlsverify},
		}
	}

	creq, err := c.Do(req)
	resp = &Response{*creq}
	return

}

func Post(url, proxy string, headers map[string]string, body io.Reader, verify bool) (*Response, error) {
	// if proxy != "" set DefaultClient.openproxy = true, else set it false
	DefaultClient.openproxy = (proxy != "")
	DefaultClient.tlsverify = verify
	if !strings.Contains(proxy, "http://") {
		DefaultClient.proxy = "http://" + proxy
	} else {
		DefaultClient.proxy = proxy
	}

	return DefaultClient.Post(url, headers, body)
}

func (c *Client) Post(endpoint string, headers map[string]string, body io.Reader) (resp *Response, err error) {
	req, err := http.NewRequest("POST", endpoint, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if c.openproxy {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: !c.tlsverify},
			Proxy:           func(r *http.Request) (*url.URL, error) { return url.Parse(c.proxy) },
		}
		c.Timeout = time.Second * 5
	} else {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: !c.tlsverify},
		}
		c.Timeout = time.Second * 3
	}

	creq, err := c.Do(req)
	resp = &Response{*creq}
	return
}
