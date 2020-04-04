package rp

import (
	"fmt"
	"github.com/just1689/plancks-proxy/config"
	"github.com/sirupsen/logrus"
	"net/http/httputil"
	"net/url"
)

func NewReverseProxy(route config.OutgoingRoute) *httputil.ReverseProxy {
	u, err := url.Parse(route.GetAddress())
	if err != nil {
		logrus.Errorln(fmt.Errorf("could not create route from url", route.GetAddress()))
		return nil
	}
	rp := httputil.NewSingleHostReverseProxy(u)
	return rp
}
