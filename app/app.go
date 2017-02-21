package app

import (
	"github.com/flyaways/storage/storage"
	"github.com/flyaways/storage/storage/adapter"
	"github.com/flyaways/tracker"

	"github.com/flyaways/storage/config"
	"github.com/flyaways/storage/util/log"
)

type App struct {
	*config.Config
	adapter.Storager
}

func New(cfg *config.Config) (*App, error) {
	var err error
	if err = log.Init(cfg.Logs.Level, cfg.Logs.Path); err != nil {
		return nil, err
	}
	app := new(App)
	app.Config = cfg
	app.Storager, err = storage.New(cfg)
	if err != nil {
		log.Error(tracker.Red("[new storage:%s]", err.Error()))
		return nil, err
	}

	return app, nil
}

func (a *App) Close() {
	log.Info(tracker.Red("[app is over]"))
}
