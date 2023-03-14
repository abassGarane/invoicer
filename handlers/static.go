package handlers

import "net/http"

func (h *DefaultHandler) ServeStatic() http.Handler {
	return http.FileServer(http.Dir("./static/"))
}
