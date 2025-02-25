package handlers

import(
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/pholophus/go_backend_practice_beginner/internal/service"
	"github.com/pholophus/go_backend_practice_beginner/internal/models"
)

type ItemHandler struct {
	Service service.ItemService
}

func NewItemHandler(s service.ItemService) *ItemHandler {
	return &ItemHandler{Service: s}
}

func (h *ItemHandler) Items(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case http.MethodGet:
		items : h.Service.GetItems()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	case http.MethodPost:
		var newItem models.Item
		err : json.NewDecoder(r.Body).Decode(&newItem)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		createdItem := h.Service.CreateItem(newItem)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdItem)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ItemHandler) Item(w http.ResponseWriter, r *http.Request) {
	// URL pattern: /items/{id}
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		item, found := h.Service.GetItem(id)
		if !found {
			http.Error(w, "Item not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(item)
	case http.MethodPut:
		var updatedItem models.Item
		if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		updatedItem.ID = id // Ensure the ID in the URL is used.
		item, ok := h.Service.UpdateItem(updatedItem)
		if !ok {
			http.Error(w, "Item not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(item)
	case http.MethodDelete:
		if ok := h.Service.DeleteItem(id); !ok {
			http.Error(w, "Item not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}