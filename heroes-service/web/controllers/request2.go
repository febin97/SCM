package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)



var a string 
	var b string 
	var c string 
	var d string 
	var f string 
	var g string 
	//var key string



func (app *Application) RequestHandler2(w http.ResponseWriter, r *http.Request) {
	blockData, err := app.Fabric.QueryHello()
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	type Test struct {
	//
	Key string `json:"key"`	
	//
	IMEINo string `json:"imeino"`
	Specifications string `json:"specifications"`
	ProducerName string `json:"producername"`
	ManufacturerName string `json:"manufacturername"`
	ManufacturingSite string `json:"manufacturingsite"`
	FinalAssemblyDate string `json:"finalassemblydate"`
	PackagingDate string `json:"packagingdate"`
	Price string `json:"price"`
	}

	type Data struct {
		Key    string `json:"key"`
		Record Test    `json:"record"`
	}

	type RecordHistory struct {
		TxId      string `json:"TxId"`
		Value     Test    `json:"Value"`
		Timestamp string `json:"Timestamp"`
		IsDelete  string `json:"IsDelete"`
	}

	var data []Data
	json.Unmarshal([]byte(blockData), &data)

	returnData := &struct {
		ResponseData         []Data
		TransactionRequested string
		TransactionUpdated   string
		RecordHistory        RecordHistory
	}{
		ResponseData:         data,
		TransactionRequested: "true",
	}
	// Query History Using Key
	
	if r.FormValue("requested") == "true" {
		// Retrieving Single Query
		QueryValue := r.FormValue("KeySearch")
		blockHistory, _ := app.Fabric.GetHistory(QueryValue)
		var queryResponse []RecordHistory
		//var queryResponse1 []RecordHistory
		
		json.Unmarshal([]byte(blockHistory), &queryResponse)
		//
		i := 0
		for i=0;i<len(queryResponse);i++ {
			queryResponse[i].Value.Key = QueryValue
		}
		//
		//queryResponse1[0] = queryResponse[len(queryResponse)-1]
		//queryResponse1[0].TxId = ""
		//queryResponse1[0].Value = queryResponse[i].Value
		//queryResponse1[0].Timestamp = ""
		//queryResponse1[0].IsDelete = ""
		//fmt.Println("lenth",len(queryResponse))
		returnData.RecordHistory = queryResponse[len(queryResponse)-1]
		returnData.TransactionRequested = "true"
		fmt.Println("### Response History ###")
		//fmt.Printf("%s", blockHistory)
		fmt.Println(blockHistory)
		fmt.Println(queryResponse)
		a = queryResponse[len(queryResponse)-1].Value.Key
		b = queryResponse[len(queryResponse)-1].Value.IMEINo
		c = queryResponse[len(queryResponse)-1].Value.Specifications
		d = queryResponse[len(queryResponse)-1].Value.ProducerName
		f = queryResponse[len(queryResponse)-1].Value.ManufacturingSite
		g = queryResponse[len(queryResponse)-1].Value.FinalAssemblyDate
	}
	
	renderTemplate2(w, r, "request2.html", returnData)
}



func (app *Application) RequestHandlerKeng(w http.ResponseWriter, r *http.Request) {
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
		key := a
		productData := Data{}
		productData.IMEINo = b
		productData.Specifications =  c
		productData.ProducerName = d
		productData.ManufacturerName = r.FormValue("manufacturername")
		productData.ManufacturingSite =  f
		productData.FinalAssemblyDate =  g
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
		http.Redirect(w, r, "/home2.html", 302)
	}
	
}
