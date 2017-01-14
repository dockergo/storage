package storage

import (
	"github.com/flyaways/tracker"
	"github.com/gin-gonic/gin"

	"github.com/flyaways/storage/app"
	"github.com/flyaways/storage/config"
	"github.com/flyaways/storage/util/log"
)

type Server struct {
	engin *gin.Engine
	app   *app.App
}

func New(cfg *config.Config) (*Server, error) {
	var err error
	s := new(Server)

	s.engin = gin.New()
	gin.SetMode(gin.ReleaseMode)

	s.app, err = app.New(cfg)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Server) Run() {
	log.Info("[begin running gin....]")

	regRouters(s.app, s.engin)

	log.Info("[start server at: %s]", tracker.Red(s.app.Config.HTTP.HTTPAddr))
	s.engin.Run(s.app.Config.HTTP.HTTPAddr)
}
