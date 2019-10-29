package config

type (
	HttpSectionMethod interface {
		GetHost() string
		GetPort() int
	}

	HttpSection struct {
		HttpSectionMethod
		Host string `yml: "host"`
		Port int `yml: "port"`
	}
)

func (section *HttpSection) GetHost () string {
	return section.Host
}

func (section *HttpSection) GetPort () int  {
	return section.Port
}