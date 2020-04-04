package handler

import (
	"github.com/just1689/plancks-proxy/config"
	"github.com/just1689/plancks-proxy/io/proxy/rp"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetHandlerHTTP(r *http.Request, c config.Config) func(w http.ResponseWriter, r *http.Request) {
	entry, err := c.FindEntry(r)
	if err != nil {
		logrus.Println("could not find route to handle")
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO:
		// Figure out how to choose route!
		route := entry.OutgoingRoutes[0]
		mp := rp.ReverseProxyHub.GetProxy(route)
		mp.ReverseProxy.ServeHTTP(w, r)
	}
}
