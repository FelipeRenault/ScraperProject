package collector

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

const (
	defaultTimeout = 2 * time.Minute
)

// Can use c.SetProxy to set a proxy server
func NewCollector() (*colly.Collector, error) {
	c := colly.NewCollector()

	c.WithTransport(&http.Transport{
		DialContext:     (&net.Dialer{KeepAlive: 5 * time.Minute}).DialContext,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	})

	c.SetRequestTimeout(defaultTimeout)

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	return c, nil
}
