package app

import (
	"github.com/flyaways/storage/agent/storage"
	"github.com/flyaways/storage/agent/storage/adapter"
	"github.com/flyaways/tracker"

	"github.com/flyaways/storage/agent/config"
	"github.com/flyaways/storage/agent/util/log"
)

type App struct {
	*config.Config
	adapter.Storager
}

func NewApp(cfg *config.Config) (*App, error) {
	var err error
	if err = log.Init(cfg.Logs.Level, cfg.Logs.Path); err != nil {
		return nil, err
	}
	app := new(App)
	app.Config = cfg
	app.Storager, err = storage.NewStorage(cfg)
	if err != nil {
		log.Error("[NewStorage:%s]", tracker.Red(err.Error()))
		return nil, err
	}

	return app, nil
}

func (a *App) Close() {
	log.Info(tracker.Red("[app is over]"))
}
