package main

import (
	"fmt"
	"net/http"

	controller "github.com/epileftro85/goapi/Controller"
	router "github.com/epileftro85/goapi/Http"
)

var (
	httpRouter     router.Router             = router.NewChiRouter()
	postController controller.PostController = controller.NewPostController()
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
