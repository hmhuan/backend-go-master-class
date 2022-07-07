package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/hmhuan/backend-go-master-class/db/sqlc"
)

// servers http request for banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// New server instance
func NewServer(store db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()

	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.getListAccount)
	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
