package routes

import (
  "fmt"
  "net/http"
  "app/facebook"
  "github.com/julienschmidt/httprouter"
)

// Define API routes
func Router() *httprouter.Router {
    router := httprouter.New()

    // Route
    router.GET("/", func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        fmt.Fprint(w, "Hi! I am ToledoBot up and running\n")
    })

    // Facebook webhook
    router.GET("/facebook/index", facebook.Index)
    router.GET("/facebook/hello/:name", facebook.Hello)

    return router
}