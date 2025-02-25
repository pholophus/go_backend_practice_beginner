package routes

import(
	"net/http"

	"github.com/pholophus/go_backend_practice_beginner/internal/handlers"
	"github.com/pholophus/go_backend_practice_beginner/internal/middleware"
	"github.com/pholophus/go_backend_practice_beginner/internal/repository"
	"github.com/pholophus/go_backend_practice_beginner/internal/service"
)

func SetupRouter() http.Handler {
	mux := http.NewServeMux()

	repo := repository.NewInMemoryItemRepository()
	svc := service.NewItemService(repo)
	itemHandler := handlers.NewItemHandler(svc)

	mux.HandleFunc("/items", middleware.BasicAuth(itemHandler.Items))
	mux.HandleFunc("/items/", middleware.BasicAuth(itemHandler.Item))

	return mux
}