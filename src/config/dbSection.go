package config

type (
	DbSectionMethod interface {
		GetPort() int
		GetHost() string
		GetUser() string
		GetBase() string
		GetPassword() string
	}

	DbSection struct {
		DbSectionMethod
		Host string `yml: "host"`
		Port int `yml: "port"`
		User string `yml: "user"`
		Pass string `yml: "pass"`
		Base string `yml: "base"`
	}
)

func (section *DbSection) GetPort() int {
	return section.Port
}

func (section *DbSection) GetUser() string {
	return section.User
}

func (section *DbSection) GetPassword() string {
	return section.Pass
}

func (section *DbSection) GetHost() string {
	return section.Host
}

func (section *DbSection) GetBase() string {
	return section.Base
}
