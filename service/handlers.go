package service

import (
	"github.com/unrolled/render"
	"net/http"
	"github.com/gorilla/mux"
)

func getFulfillmentStatusHandler(formatter *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		sku := vars["sku"]
		formatter.JSON(writer, http.StatusOK, fulfillmentStatus{
			SKU:             sku,
			ShipsWithin:     14,
			QuantityInStock: 100,
		})
	}
}

func rootHandler(formatter *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		formatter.Text(writer, http.StatusOK, "Fulfillment Service, see http://github.com/cloudnativego/backing-fulfillment for API.")
	}
}