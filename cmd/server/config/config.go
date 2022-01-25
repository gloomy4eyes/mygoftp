package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

type Config struct {
	host       string
	port       string
	storageDNS string
}

func NewFromFile(configFilePath string) *Config {
	k := koanf.New(".")
	k.Load(file.Provider(configFilePath), yaml.Parser())

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
