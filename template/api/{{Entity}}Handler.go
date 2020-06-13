package api
import (
	domain "{{Module}}"
	"net/http"
)
type {{Entity}}Handler struct {
	service domain.{{Entity}}Service
}
func (h *{{Entity}}Handler) index(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("It Works"))
}
