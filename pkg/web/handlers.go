package web

import (
	"github.com/bmerchant22/pkg/store"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	r     *gin.Engine
	store *store.MongoStore
}

func (srv *Server) RecentActionsHandler(ctx *gin.Context) {
	recentActions, err := srv.store.QueryRecentActions()
	if err != nil {
		log.Printf("Error occurred while fetching recentActions: %v", err)
		ctx.String(http.StatusBadRequest, "Error while getting recent actions")
	}

	ctx.JSON(http.StatusOK, recentActions)
}

func (srv *Server) StartListeningRequests(addr string) error {
	err := srv.r.Run(addr)
	if err != nil {
		log.Printf("Error occurred while starting server: %v", err)
		return nil
	}
	return err
}
