package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/Zanda256/gcloud-workflow-poc/api/models"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
)

var db map[string]*models.MainInput

var callbackDb map[string]*models.QcWfCallback

// GET /orders/:order/docs

func GetDocStatus(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	st := models.Status{}
	oId := params.ByName("order")
	for k, v := range db {
		if k == oId {
			st.DocStatus = v.State.DocStatus
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(st); err != nil {
				log.Fatalf("failed to encode json %+v", st)
			}
			return
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	log.Printf("resource with id %s not found", oId)
	return
}

// POST /orders/:order/qc/callback

func StoreCallback(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var cb models.QcWfCallback
	id := params.ByName("order")
	m, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("failed to read req body")
	}
	err = json.Unmarshal(m, &cb)
	if err != nil {
		log.Fatalf("failed to unmarshal callback json")
	}
	callbackDb[id] = &cb
	w.WriteHeader(http.StatusOK)
}

// POST /orders/:order/qc

func UpdateQc(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("order")
	cb := callbackDb[id]
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("failed to read req body")
	}
	callWorkflow(data, cb.Url)
	w.WriteHeader(http.StatusOK)
	return
}

func callWorkflow(data []byte, url string){
	c := http.DefaultClient
	body := bytes.NewReader(data)
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		log.Fatalf("failed to create http request")
	}
	res, err := c.Do(req)
	if err != nil {
		log.Fatalf("failed to call callback")
	}
	log.Printf("Call back response: %d\n", res.StatusCode)
}

//var stat models.Status
//err = json.Unmarshal(data, &stat)
//if err != nil {
//log.Fatalf("failed to unmarshal qc json")
//}