package handler

import "net/http"

type PageHandler struct {
	Renderer *Renderer
}

func NewPageHandler(renderer *Renderer) *PageHandler {
	return &PageHandler{
		Renderer: renderer,
	}
}

func (h *PageHandler) Home(w http.ResponseWriter, r *http.Request) {
	h.Renderer.Render(w, "home.html", nil)
}
