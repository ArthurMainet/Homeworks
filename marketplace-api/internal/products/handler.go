package products

import (
	"Email-API/config"
	"Email-API/packages/api"
	"Email-API/packages/middlewares"
	"Email-API/packages/responce"
	"net/http"
	"strconv"
)

type ProductHandlerDeps struct {
	Config            *config.AuthConfig
	ProductRepository *ProductRepository
}

type ProductHandler struct {
	ProductRepository *ProductRepository
	Config            *config.AuthConfig
}

func NewProductHandler(router *http.ServeMux, deps ProductHandlerDeps) {
	prHandler := &ProductHandler{
		ProductRepository: deps.ProductRepository,
		Config:            deps.Config,
	}
	router.HandleFunc("GET /products/{id}", prHandler.GetbyID())
	router.HandleFunc("POST /products", prHandler.Create())
	router.HandleFunc("PUT /products/{id}", prHandler.fullUpdate())
	router.HandleFunc("PATCH /products/{id}", prHandler.Update())
	router.Handle("DELETE /products/{id}", middlewares.IsAuth(prHandler.Delete(), prHandler.Config))
}

func (handler *ProductHandler) GetbyID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Invalid ID.", http.StatusBadRequest)
			return
		}
		product, err := handler.ProductRepository.GetByID(id)
		if err != nil {
			http.Error(w, "No products with this ID.", http.StatusNotFound)
			return
		}
		responce.ResponceJSON(w, product, http.StatusOK)
	}
}

func (handler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := api.HandleReq[ProductRequest](w, r)
		if err != nil {
			http.Error(w, "Invalid form.", http.StatusBadRequest)
			return
		}
		product := NewProduct(body)
		_, err = handler.ProductRepository.Create(product)
		if err != nil {
			http.Error(w, "Didn't create. Retry please.", http.StatusBadGateway)
			return
		}
		responce.ResponceJSON(w, "Product created successfully", 200)
	}
}

func (handler *ProductHandler) fullUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Invalid ID.", http.StatusBadRequest)
			return
		}

		product, err := handler.ProductRepository.GetByID(id)
		if err != nil {
			http.Error(w, "No data with this ID.", http.StatusNotFound)
			return
		}

		body, err := api.HandleReq[ProductRequest](w, r)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		product.Name = body.Name
		product.Description = body.Description
		product.Images = body.Images
		product.Price = body.Price

		updatedProduct, err := handler.ProductRepository.Update(product)
		if err != nil {
			http.Error(w, "Update error", http.StatusInternalServerError)
			return
		}
		responce.ResponceJSON(w, updatedProduct, 200)
	}
}

func (handler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Invalid ID.", http.StatusBadRequest)
			return
		}

		body, err := api.HandleReq[UpdateProductRequest](w, r)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		product, err := handler.ProductRepository.GetByID(id)
		if err != nil {
			http.Error(w, "No data with this ID.", http.StatusBadRequest)
			return
		}

		if body.Name != nil {
			product.Name = *body.Name
		}
		if body.Description != nil {
			product.Description = *body.Description
		}
		if body.Images != nil {
			product.Images = *body.Images
		}
		if body.Price != nil {
			product.Price = *body.Price
		}

		updatedProduct, err := handler.ProductRepository.Update(product)
		if err != nil {
			http.Error(w, "Update error", http.StatusInternalServerError)
			return
		}
		responce.ResponceJSON(w, updatedProduct, 200)
	}
}

func (handler *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Invalid ID.", http.StatusBadRequest)
			return
		}

		_, err = handler.ProductRepository.GetByID(id)
		if err != nil {
			http.Error(w, "Delete error. No product with this ID.", http.StatusBadRequest)
			return
		}

		err = handler.ProductRepository.Delete(id)
		if err != nil {
			http.Error(w, "Delete error.", http.StatusInternalServerError)
			return
		}

		responce.ResponceJSON(w, "Product successfully delete.", http.StatusOK)
	}
}
