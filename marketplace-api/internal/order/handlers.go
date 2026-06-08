package order

import (
	"Email-API/packages/api"
	"Email-API/packages/responce"
	"log"
	"net/http"
	"strconv"
)

type OrderHandlersDeps struct {
	OrderService *OrderService
}

type OrderHandlers struct {
	OrderService *OrderService
}

func NewOrderHandlers(router *http.ServeMux, deps OrderHandlersDeps) {
	orderHandlers := &OrderHandlers{
		OrderService: deps.OrderService,
	}

	router.HandleFunc("POST /order", orderHandlers.CreateOrder())
	router.HandleFunc("GET /order/{id}", orderHandlers.GetOrderByID())
	router.HandleFunc("GET /my-orders", orderHandlers.GetUserOrders())
}

func (handler *OrderHandlers) CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := api.DecodeJSON[OrderRequest](r.Body)
		if err != nil {
			http.Error(w, "invalid form", http.StatusBadRequest)
			return
		}

		cookie, err := r.Cookie("userid")
		if err != nil {
			http.Error(w, "cookie err", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		userId, _ := strconv.Atoi(cookie.Value)

		products, err := handler.OrderService.GetProducts(body.ProductIDs)
		if err != nil {
			http.Error(w, "DB err", http.StatusBadRequest)
			log.Println(err)
			return
		}
		order := NewOrder(uint(userId), products)

		_, err = handler.OrderService.Create(order)
		if err != nil {
			http.Error(w, "DB err", http.StatusBadRequest)
			log.Println(err)
			return
		}

		responce.ResponceJSON(w, "Order created", http.StatusOK)
	}
}

func (handler *OrderHandlers) GetOrderByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "ivalid form", http.StatusBadRequest)
			log.Println(err)
			return
		}
		id := uint(path)

		order, err := handler.OrderService.Repo.GetOrderById(id)
		if err != nil {
			http.Error(w, "DB err", http.StatusBadRequest)
			log.Println(err)
			return
		}

		responce.ResponceJSON(w, order, http.StatusOK)
	}
}

func (handler *OrderHandlers) GetUserOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("userid")
		if err != nil {
			http.Error(w, "cookie err", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		userId, _ := strconv.Atoi(cookie.Value)

		order, err := handler.OrderService.Repo.GetOrderByUserId(uint(userId))
		if err != nil {
			http.Error(w, "DB err", http.StatusBadRequest)
			log.Println(err)
			return
		}

		responce.ResponceJSON(w, order, http.StatusOK)
	}
}
