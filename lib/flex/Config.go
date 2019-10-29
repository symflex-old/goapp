package flex

type (
	Config interface {
		Load(filePath string) *Config
		Parse()
	}
)

func (config *Config) Load(filePath string) *Config {
	return config
}