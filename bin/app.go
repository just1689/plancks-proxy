package main

import (
	"flag"
	"fmt"
	"github.com/just1689/plancks-proxy/config"
	"github.com/just1689/plancks-proxy/io/disk"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"time"
)

var generateConfig = flag.String("generate", "", "specify a filename to generate config. Terminates.")

func main() {
	flag.Parse()
	if *generateConfig != "" {
		generateExampleConfig(*generateConfig)
		return
	}

	configFile := os.Getenv("config")
	if configFile == "" {
		logrus.Errorln("config env var is empty. Nothing left to do.")
		return
	}

	c := disk.WatchFile(configFile, 2*time.Second)
	go func() {
		for b := range c {
			fmt.Println(string(b))
		}
	}()

	for {
		time.Sleep(5 * time.Second)
	}

}

func generateExampleConfig(f string) {
	example := &config.Config{
		Entries: []config.Entry{
			{
				AllowHTTP:    true,
				StripContext: false,
				IncomingRoutes: []config.IncomingRoute{
					{
						IncludeHosts: []string{"example.com"},
						Context:      "",
						ProvisionTLS: config.ProvisionTLS{
							LetsEncryptEmail:       "",
							LetsEncryptAcceptTerms: false,
						},
					},
				},
				OutgoingRoutes: []config.OutgoingRoute{
					{
						Address:    "192.168.1.10",
						ConnectTLS: false,
					},
				},
			},
		},
		WriteTimeout: 60,
		ReadTimeout:  60,
		HTTPAddress:  ":80",
		TLSAddress:   ":443",
	}
	d, err := yaml.Marshal(example)
	if err != nil {
		logrus.Fatalf("error: %v", err)
	}
	err = ioutil.WriteFile(f, d, 0644)
	if err != nil {
		logrus.Fatalln(err)
	}

}

//func getFromEnv() config.Config {
//
//}
