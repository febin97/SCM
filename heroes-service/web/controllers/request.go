package controllers

import (
	"net/http"
	"encoding/json"
)
type Data struct {
	IMEINo   string `json:"imeino"`
	/*Processor string `json:"processor"`
	Dimensions string `json:"dimensions"`
	Battery string `json:"battery"`
	DisplayUnit string `json:"displayunit"`
	CameraModule string `json:"cameramodule"` 
	Memory string `json:"memory"`*/
	Specifications string `json:"specifications"`
	ProducerName string `json:"producername"`
	ManufacturerName string `json:"manufacturername"`
	ShelfLife string `json:"shelflife"`
	ManufacturingSite string`json:"manufacturingsite"`
	FinalAssemblyDate string `json:"finalassemblydate"`
	PackagingDate string `json:"packagingdate"`
	Price string `json:"price"`
}

func (app *Application) RequestHandler(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		TransactionId string
		Success       bool
		Response      bool
	}{
		TransactionId: "",
		Success:       false,
		Response:      false,
	}
	if r.FormValue("submitted") == "true" {
		key := r.FormValue("key")
		productData := Data{}
		productData.IMEINo = r.FormValue("imeino")
		productData.Specifications = r.FormValue("specifications")
		productData.ProducerName = r.FormValue("producername")
		productData.ManufacturerName = r.FormValue("manufacturername")
		productData.ShelfLife = r.FormValue("shelflife")
		productData.ManufacturingSite = r.FormValue("manufacturingsite")
		productData.FinalAssemblyDate = r.FormValue("finalassemblydate")
		productData.PackagingDate = r.FormValue("packagingdate")
		productData.Price = r.FormValue("price")
        RequestData, _ := json.Marshal(productData)
		txid, err := app.Fabric.InvokeHello(key,string(RequestData))
		if err != nil {
			http.Error(w, "Unable to invoke hello in the blockchain", 500)
		}
		data.TransactionId = txid
		data.Success = true
		data.Response = true
	}
	renderTemplate2(w, r, "request.html", data)
}
