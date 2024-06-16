package httpServer

import "net/http"

type CorsMiddleware struct {
	handler http.Handler
}

func (cm CorsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	cm.handler.ServeHTTP(w, r)
}

func UseCorsMiddleWare(handler http.Handler) CorsMiddleware {
	return CorsMiddleware{handler}
}