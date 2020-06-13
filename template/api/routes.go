package api
import (
	"github.com/urfave/negroni"
	"net/http"
)
func Routes(mux *http.ServeMux) {
	n := negroni.Classic()
	n.UseHandler(mux)
	//Routes
	handler := {{Entity}}Handler{}
	mux.HandleFunc("/", handler.index)
}
