package modHttp

import (
	"datacontrol/modUtility"
	"fmt"

	"github.com/gin-gonic/gin"
)

type CHttpServer struct {
	hServer *gin.Engine
}

var g_singleHttpServer *CHttpServer = &CHttpServer{}

func getSingleHttpServer() *CHttpServer {
	return g_singleHttpServer
}

func (pSelf *CHttpServer) initialize() error {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	pSelf.hServer = r
	return nil
}

func (pSelf *CHttpServer) start() error {
	return pSelf.hServer.Run(fmt.Sprintf(":%d", modUtility.G_HttpPort))
}
