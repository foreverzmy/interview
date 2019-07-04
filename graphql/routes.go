package main

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

// Route 路由
type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

// GPHandle graphql http handle
func GPHandle(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	// claims, err := auth.CheckToken(r)
	// if err != nil {
	// 	auth.ResponseWithJSON(w, http.StatusUnauthorized, err.Error())
	// } else {
	// 	// ctx = context.WithValue(ctx, auth.TokenUserIDKey, claims.UserID)
	// }
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

	h.ContextHandler(ctx, w, r)
}

func init() {
	// register("POST", "/auth/register", auth.Register, nil)
	// register("POST", "/auth/login", auth.Login, nil)
	register("POST", "/graphql", GPHandle, nil)

	register("OPTIONS", "/graphql", GPHandle, nil)
	// register("POST", "/gp", controller.GPHandle, auth.TokenMiddleware)
}

// NewRouter 新建路.
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		r := router.Methods(route.Method).
			Path(route.Pattern)
		if route.Middleware != nil {
			r.Handler(route.Middleware(route.Handler))
		} else {
			r.Handler(route.Handler)
		}
	}
	return router
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
