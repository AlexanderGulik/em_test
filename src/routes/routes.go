package routes

import (
	"em_test/src/subscription/controller"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) {
	
		subscription := controller.NewSubController()

		mux.HandleFunc("POST /created-sub", subscription.CreateSub)
		mux.HandleFunc("GET /get-sub/{id}", subscription.GetSubId)
		mux.HandleFunc("GET /get-sub", subscription.SelectSubAll)
		mux.HandleFunc("DELETE /delete-sub/{id}", subscription.DeleteSub)
		mux.HandleFunc("PUT /update-sub/{id}", subscription.UpdateSubFull)
		mux.HandleFunc("PATCH /update-sub/{id}", subscription.UpdateSubPartial)
		mux.HandleFunc("GET /get-sub-sum",subscription.SelectGetSumSub)

}
