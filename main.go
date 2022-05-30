package main

import (
	"github.com/Zanda256/gcloud-workflow-poc/api/handlers"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/orders/:order/docs", handlers.GetDocStatus)
	router.POST("/orders/:order/qc/callback", handlers.StoreCallback)
	router.POST("/orders/:order/qc", handlers.UpdateQc)

	log.Fatal(http.ListenAndServe(":8080", router))
}
