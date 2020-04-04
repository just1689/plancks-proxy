package rp

import (
	"github.com/just1689/plancks-proxy/config"
	"net/http/httputil"
	"sync"
)

var ReverseProxyHub = &Hub{
	proxies: make(map[string]*MonitoredProxy),
}

type Hub struct {
	sync.Mutex
	proxies map[string]*MonitoredProxy
}

func (h *Hub) GetProxy(route config.OutgoingRoute) (mp *MonitoredProxy) {
	var found bool
	h.Lock()
	defer h.Unlock()
	mp, found = h.proxies[route.GetAddress()]
	if found {
		return
	}
	mp = &MonitoredProxy{
		healthy:      true,
		ReverseProxy: NewReverseProxy(route),
	}
	h.proxies[route.GetAddress()] = mp
	return
}

type MonitoredProxy struct {
	sync.Mutex
	healthy      bool
	ReverseProxy *httputil.ReverseProxy
}
