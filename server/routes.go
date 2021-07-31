package server

import "github.com/gin-gonic/gin"

// GET "/healthz"
func (s *server) GetHealthz(c *gin.Context) {
	c.AbortWithStatus(200)
}

// GET "/block/{block_param}"
func (s *server) GetBlock(c *gin.Context) {
	c.AbortWithStatus(200)
}

// GET "/block/{block_param}/txs/{txs_param}"
func (s *server) GetTransaction(c *gin.Context) {
	c.AbortWithStatus(200)
}
