package config

type IncomingRoute struct {
	IncludeHosts []string     `yaml:"includeHosts"`
	Context      string       `yaml:"context"`
	ProvisionTLS ProvisionTLS `yaml:"provisionTLS"`
}

type ProvisionTLS struct {
	LetsEncryptEmail       string `yaml:"letsEncryptEmail"`
	LetsEncryptAcceptTerms bool   `yaml:"letsEncryptAcceptTerms"`
}
