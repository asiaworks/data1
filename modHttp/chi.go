package modHttp

import (
	"datacontrol/modUtility"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type CHttpServer struct {
	chiInst *chi.Mux
}

var g_singleHttpServer *CHttpServer = &CHttpServer{}

func getSingleHttpServer() *CHttpServer {
	return g_singleHttpServer
}

func (pSelf *CHttpServer) initialize() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	pSelf.chiInst = r

	return nil
}

func (pSelf *CHttpServer) start() error {
	return http.ListenAndServe(fmt.Sprintf("[::]:%d", modUtility.G_HttpPort), pSelf.chiInst)
}

/*
func Chi_() {
	http.ListenAndServe(":3000", r)
}*/
