package main
/*
import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

type (
	ConfigMethod interface {
		Load(file string)
		Parse(data []byte) error
		isDebug() bool
	}

	HttpSectionMethod interface {
		GetHost() string
		GetPort() int
	}

	HttpSection struct {
		HttpSectionMethod
		Host string `yml: "host"`
		Port int `yml: "port"`
	}

	ConfigYamlMethod interface {
		GetHttp() *HttpSection
		GetHost() string
		GetPort() int
	}

	ConfigYaml struct {
		ConfigMethod
		ConfigYamlMethod
		Debug bool `yml: "debug"`
		Http HttpSection `yml: "http"`
	}
)

// Load config file
func (config *ConfigYaml) load(file string)  {
	filename, _ := filepath.Abs(file)
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("Load config: %s", err)
	}

	err = config.parse(data)

	if err != nil {
		log.Fatalf("Parse config: %s", err)
	}
}

// Parse config data
func (config *ConfigYaml) parse(data []byte) error {
	return yaml.Unmarshal(data, &config)
}

func (config *ConfigYaml) GetHttp() *HttpSection {
	return  &config.Http
}

// Get host from config
func (config *ConfigYaml) GetHost() (host string) {
	host = config.Http.Host
	return host
}

// Get port from config
func (config *ConfigYaml) GetPort() int {
	return config.GetHttp().GetPort()
}

func (config *ConfigYaml) IsDebug() bool {
	return config.Debug
}

func (config *HttpSection) GetHost () string {
	return config.Host
}

func (config *HttpSection) GetPort () int  {
	return config.Port
}*/