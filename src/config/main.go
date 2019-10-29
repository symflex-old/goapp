package config

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

	ConfigYamlMethod interface {
		GetHttp() *HttpSection
		GetDb() *DbSection
	}

	ConfigYaml struct {
		ConfigMethod
		ConfigYamlMethod
		Debug bool `yml: "debug"`
		Http *HttpSection `yml: "http"`
		Db *DbSection `yml: "db"`
	}
)

// Load config file
func (config *ConfigYaml) Load(file string)  {
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
	return  config.Http
}

func (config *ConfigYaml) GetDb() *DbSection {
	return config.Db
}

func (config *ConfigYaml) IsDebug() bool {
	return config.Debug
}
