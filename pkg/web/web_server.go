package web

import (
	"github.com/bmerchant22/pkg/store"
	"github.com/gin-gonic/gin"
)

func CreateWebServer(mongoStore *store.MongoStore) *Server {
	srv := new(Server)
	srv.r = gin.Default()
	srv.store = mongoStore

	srv.r.GET(kGetRecentActions, srv.RecentActionsHandler)

	return srv
}
