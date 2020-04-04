package v2

import (
	"github.com/just1689/plancks-proxy/config"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"time"
)

func StartupServers(config config.Config) {
	h := &HTTPServer{
		config: config,
	}
	go h.Listen()

	t := &TLSServer{
		config: config,
	}
	go t.Listen()

}

type HTTPServer struct {
	config config.Config
}

func (h *HTTPServer) Listen() {
	f := func(ctx *fasthttp.RequestCtx) {
		//TODO: implement reverse proxy here
	}
	s := fasthttp.Server{
		Handler:      f,
		WriteTimeout: time.Duration(h.config.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(h.config.ReadTimeout) * time.Second,
	}
	logrus.Panic(s.ListenAndServe(h.config.HTTPAddress))

}

type TLSServer struct {
	config config.Config
}

func (h *TLSServer) Listen() {
	//f := func(ctx *fasthttp.RequestCtx) {
	//	//TODO: implement reverse proxy here
	//}
	//s := fasthttp.Server{
	//	Handler:      f,
	//	WriteTimeout: time.Duration(h.config.WriteTimeout) * time.Second,
	//	ReadTimeout:  time.Duration(h.config.ReadTimeout) * time.Second,
	//}

	//listenTLS, err := certmagic.Listen(hosts)

}
