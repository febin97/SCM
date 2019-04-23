package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	helloValue, err := app.Fabric.QueryHello()
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	type HelloData struct {
		Key    string `json:"key"`
		Record struct{
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
	ManufacturingSite string`json:"manufacturingsite"`
	FinalAssemblyDate string `json:"finalassemblydate"`
	PackagingDate string `json:"packagingdate"`
	Price string `json:"price"`
	}
   }
	var data []HelloData
	json.Unmarshal([]byte(helloValue), &data)
	fmt.Println(data)

	returnData := &struct {
		ResponseData []HelloData
	}{
		ResponseData: data,
	}

	renderTemplate1(w, r, "home.html", returnData)
}
