package config

import "fmt"

type OutgoingRoute struct {
	Address    string
	ConnectTLS bool
}

func (o *OutgoingRoute) GetAddress() string {
	if o.ConnectTLS {
		return fmt.Sprint("http://s", o.Address)
	}
	return fmt.Sprint("http://", o.Address)
}
