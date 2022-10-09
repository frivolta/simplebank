package api

import (
	"github.com/gin-gonic/gin"
	"os"
	db "simplebank/db/sqlc"
	"testing"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	server := NewServer(store)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
