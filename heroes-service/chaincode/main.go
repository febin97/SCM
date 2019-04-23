package main

import (
	"fmt"
	"strconv"
	"encoding/json"
	"bytes"
	"time"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// HeroesServiceChaincode implementation of Chaincode
type HeroesServiceChaincode struct {
}

type Test struct {

	IMEINo   string `json:"imeino"`
	Specifications string `json:"specifications"`
	ProducerName string `json:"producername"`
	ManufacturerName string `json:"manufacturername"`
	ManufacturingSite string`json:"manufacturingsite"`
	FinalAssemblyDate string `json:"finalassemblydate"`
	PackagingDate string `json:"packagingdate"`
	Price string `json:"price"`
}



// Init of the chaincode
// This function is called only one when the chaincode is instantiated.
// So the goal is to prepare the ledger to handle future requests.
func (t *HeroesServiceChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### HeroesServiceChaincode Init ###########")

	// Get the function and arguments from the request
	function, _ := stub.GetFunctionAndParameters()

	// Check if the request is the init function
	if function != "init" {
		return shim.Error("Unknown function call")
	}

	test := []Test{
		Test{IMEINo:"mlz12345678",Specifications:"ModelName:Lenovo K3 ,Processor:Snapdragon-630,Battery: 3200, DisplayUnit:1920*720,Memory:16Gb",ProducerName:"Saleel",ManufacturerName:"Shilu",ManufacturingSite:"Hyderabad",FinalAssemblyDate:"02-01-2019",PackagingDate:"09/01/2019",Price:"Rs.8000"},

		Test{IMEINo:"mle17345678",Specifications:"ModelName:MotoG5s,Processor:Snapdragon-580,Battery:2300,DisplayUnit:1280*720,Memory:8Gb",ProducerName:"Saleel",ManufacturerName:"Sooraj",ManufacturingSite:"Bangalore",FinalAssemblyDate:"02/01/2019",PackagingDate:"08/01/2019",Price:"Rs.9000"},

		Test{IMEINo:"mlz23451235",Specifications:"ModelName: iPhone X,Processor:A11, Battery:3500,DisplayUnit:1920*720,Memory:64Gb",ProducerName:"Aby",ManufacturerName:"Shilu",ManufacturingSite:"Delhi",FinalAssemblyDate:"04/01/2019",PackagingDate:"07/01/2019",Price:"Rs.50000"}}
	i := 0
	for i < len(test) {
		fmt.Println("i is ", i)
		testAsBytes, _ := json.Marshal(test[i])
		stub.PutState("Item"+strconv.Itoa(i),testAsBytes)//[]byte{test[i]})
		fmt.Println("Added", test[i])
		i = i + 1
	}

	// Return a successful message
	return shim.Success(nil)
}

// Invoke
// All future requests named invoke will arrive here.
func (t *HeroesServiceChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### HeroesServiceChaincode Invoke ###########")

	// Get the function and arguments from the request
	function, args := stub.GetFunctionAndParameters()

	// Check whether it is an invoke request
	if function != "invoke" {
		return shim.Error("Unknown function call")
	}

	// Check whether the number of arguments is sufficient
	if len(args) < 1 {
		return shim.Error("The number of arguments is insufficient.")
	}

	// In order to manage multiple type of request, we will check the first argument.
	// Here we have one possible argument: query (every query request will read in the ledger without modification)
	if args[0] == "query" {
		return t.query(stub, args)
	}

	// The update argument will manage all update in the ledger
	if args[0] == "invoke" {
		return t.invoke(stub, args)
	}

	if args[0] == "gethistory" {
		return t.gethistory(stub, args)
	}

	// If the arguments given don’t match any function, we return an error
	return shim.Error("Unknown action, check the first argument")
}

// query
// Every readonly functions in the ledger will be here
func (t *HeroesServiceChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### HeroesServiceChaincode query ###########")

	// Check whether the number of arguments is sufficient
	if len(args) < 2 {
		return shim.Error("The number of arguments is insufficient.")
	}

	// Like the Invoke function, we manage multiple type of query requests with the second argument.
	// We also have only one possible argument: hello
	if args[1] == "hello" {

		// GetState by passing lower and upper limits
		resultsIterator, err := stub.GetStateByRange("", "")
		if err != nil {
			return shim.Error(err.Error())
		}
		defer resultsIterator.Close()

		// buffer is a JSON array containing QueryResults
		var buffer bytes.Buffer
		buffer.WriteString("[")

		bArrayMemberAlreadyWritten := false
		for resultsIterator.HasNext() {
			queryResponse, err := resultsIterator.Next()
			if err != nil {
				return shim.Error(err.Error())
			}
			// Add a comma before array members, suppress it for the first array member
			if bArrayMemberAlreadyWritten == true {
				buffer.WriteString(",")
			}
			buffer.WriteString("{\"Key\":")
			buffer.WriteString("\"")
			buffer.WriteString(queryResponse.Key)
			buffer.WriteString("\"")

			buffer.WriteString(", \"Record\":")
			// Record is a JSON object, so we write as-is
			buffer.WriteString(string(queryResponse.Value))
			//fmt.Println(string(queryResponse.Value)
			buffer.WriteString("}")
			bArrayMemberAlreadyWritten = true
		}
		buffer.WriteString("]")

		fmt.Printf("- queryAll:\n%s\n", buffer.String())

		return shim.Success(buffer.Bytes())
	}

	// If the arguments given don’t match any function, we return an error
	return shim.Error("Unknown query action, check the second argument.")
}

// invoke
// Every functions that read and write in the ledger will be here
func (t *HeroesServiceChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### HeroesServiceChaincode invoke ###########")

	if len(args) < 3 {
		return shim.Error("The number of arguments is insufficient.")
	}else {

		// Write the new value in the ledger
		var newProduct Test
		//newProduct= []byte{args[2]}
	    json.Unmarshal([]byte(args[2]), &newProduct)
		fmt.Printf("%s",newProduct.Specifications)
		var test = Test{IMEINo:newProduct.IMEINo,Specifications:newProduct.Specifications,ProducerName:newProduct.ProducerName,ManufacturerName:newProduct.ManufacturerName,ManufacturingSite:newProduct.ManufacturingSite,FinalAssemblyDate:newProduct.FinalAssemblyDate,PackagingDate:newProduct.PackagingDate,Price:newProduct.Price}
		testAsBytes, _ := json.Marshal(test)
		err:=stub.PutState(args[1],testAsBytes)
		err = stub.SetEvent("eventUpdateRecords", []byte{})
		//err := stub.PutState(args[1], []byte(args[2]))
		if err != nil {
			return shim.Error("Failed to update state of hello")
		}

		// Notify listeners that an event "eventInvoke" have been executed (check line 19 in the file invoke.go)
		err = stub.SetEvent("eventInvoke", []byte{})
		if err != nil {
			return shim.Error(err.Error())
		}

		// Return this value in response
		return shim.Success(nil)
	}

	// If the arguments given don’t match any function, we return an error
	return shim.Error("Unknown invoke action, check the second argument.")
}

func (s *HeroesServiceChaincode) gethistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
     
	Key := args[1]
	

	resultsIterator, err := stub.GetHistoryForKey(Key)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForMarble returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func main() {
	// Start the chaincode and make it ready for futures requests
	err := shim.Start(new(HeroesServiceChaincode))
	if err != nil {
		fmt.Printf("Error starting Heroes Service chaincode: %s", err)
	}
}
