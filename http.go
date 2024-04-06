package kadigo

import (
	"net/http"

	"github.com/go-chi/render"
)

func Render(w http.ResponseWriter, r *http.Request, body any) {
	render.JSON(w, r, body)
}
