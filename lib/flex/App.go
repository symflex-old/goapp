package flex

type (
	AppMethod interface {
		Bootstrap(config *Config) *App
		Process() *App
	}

	App struct {
		AppMethod
		Service []interface{}
	}
)

func (app *App) Bootstrap(config *Config) *App {
	return  app
}

func (app *App) Process() *App {
	return  app
}