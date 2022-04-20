package main

import (
	"fmt"
	"net/http"

	controller "github.com/epileftro85/goapi/Controller"
	router "github.com/epileftro85/goapi/Http"
	repository "github.com/epileftro85/goapi/Repository"
	service "github.com/epileftro85/goapi/Service"
)

var (
	httpRouter     router.Router             = router.NewChiRouter()
	postRepos      repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepos)
	postController controller.PostController = controller.NewPostController(postService)
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
