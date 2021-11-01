package config

type Config struct {
	host       string
	port       string
	storageDNS string
}

func NewFromFile(configFilePath string) *Config {
	return &Config{}
}

func (c Config) Host() string {
	return c.host
}

func (c Config) Port() string {
	return c.port
}

func (c Config) StorageDNS() string {
	return c.storageDNS
}
