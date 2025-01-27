package child

import (
	"net/http"

	"github.com/gorilla/mux"
)

type MyHandler struct {
	// MyHandler fields
	Router *mux.Router
	Http *http.Server
}