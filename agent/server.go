package agent

import (
	"github.com/Flyaways/tracker"
	"github.com/gin-gonic/gin"

	"github.com/flyaways/storage/agent/app"
	"github.com/flyaways/storage/agent/config"
	"github.com/flyaways/storage/agent/util/log"
)

type Server struct {
	engin *gin.Engine
	app   *app.App
}

func NewServer(cfg *config.Config) (*Server, error) {
	var err error
	s := new(Server)

	s.engin = gin.New()
	gin.SetMode(gin.ReleaseMode)

	s.app, err = app.NewApp(cfg)
	if err != nil {
		return nil, err
	}

	log.Info("[initialize waitting....]")
	StorageInit(s.app)
	log.Info("[initialize pass!]")
	return s, nil
}

func (s *Server) Run() {
	log.Info("[begin running gin....]")

	s.prepare()

	log.Info("[start server at: %s]", tracker.Red(s.app.Config.HTTP.HTTPAddr))
	s.engin.Run(s.app.Config.HTTP.HTTPAddr)
}

func (s *Server) prepare() {

	RegisterURLs(s.app, s.engin)

}
