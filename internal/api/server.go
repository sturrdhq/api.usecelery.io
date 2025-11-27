package api

import (
	"fmt"
	"github.com/sturrdhq/celery-server/internal/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type Server struct {
	Port         int
	db           *database.DBClient
	routerEngine *gin.Engine
}

// NewServer creates a new server with the default values
func NewServer(port int, db *database.DBClient) *Server {
	routerEngine := gin.Default()

	return &Server{
		port,
		db,
		routerEngine,
	}
}

// Start sets up all the routes and starts the server
func (s *Server) Start(port int) {
	address := fmt.Sprintf(":%d", port)
	log.Fatal(s.routerEngine.Run(address))
}
