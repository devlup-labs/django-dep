package router

import (
	"github.com/devlup-labs/django-dep/handler"
	"github.com/devlup-labs/django-dep/types"
	"github.com/emicklei/go-restful"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	wsContainer := restful.NewContainer()
	wsContainer.EnableContentEncoding(true)

	apiV1Ws := new(restful.WebService)
	apiV1Ws.Path("").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	apiV1Ws.Route(
		apiV1Ws.POST("/").
			To(handler.Deploy).
			Reads(types.RequestPayload{}),
	)

	wsContainer.Add(apiV1Ws)

	r.Methods(http.MethodGet).Path("/ping").HandlerFunc(handler.Ping)
	r.PathPrefix("/").Handler(wsContainer)

	return r
}
