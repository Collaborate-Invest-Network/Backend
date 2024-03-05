package router

import (
	"github.com/Collaborate-Invest-Network/Backend/internal/handler"
	"github.com/gorilla/mux"
)

// APIRouter is the main router
var APIRouter = mux.NewRouter()

func init() {

	APIRouter.Get("/").HandlerFunc(handler.Home)

}
