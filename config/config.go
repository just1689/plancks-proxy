package config

import (
	"errors"
	"net/http"
)

type Config struct {
	Entries      []Entry `yaml:"entries"`
	WriteTimeout int     `yaml:"writeTimeout"`
	ReadTimeout  int     `yaml:"readTimeout"`
	HTTPAddress  string  `yaml:"HTTPAddress"`
	TLSAddress   string  `yaml:"TLSAddress"`
}

func (c Config) FindEntry(r *http.Request) (entry Entry, err error) {
	for _, entry = range c.Entries {
		for _, i := range entry.IncomingRoutes {
			for _, host := range i.IncludeHosts {
				if r.Host == host {
					if i.Context == "" {
						return
					}
					//TODO: check that is this is right
					if r.RequestURI == i.Context {
						return
					}
				}
			}
		}
	}
	err = errors.New("could not find route")
	return
}

type Entry struct {
	AllowHTTP      bool            `yaml:"allowHTTP"`
	StripContext   bool            `yaml:"stripContext"`
	IncomingRoutes []IncomingRoute `yaml:"incomingRoutes"`
	OutgoingRoutes []OutgoingRoute `yaml:"outgoingRoutes"`
}
