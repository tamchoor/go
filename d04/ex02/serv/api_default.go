/*
 * Candy Server
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

//#include "cow.h"
import "C"

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"unsafe"
)

const (
	CE = 10
	AA = 15
	NT = 17
	DE = 21
	YR = 23
)

func countChange(order Order) int32 {
	var price int32 = 0
	if order.CandyType == "CE" {
		price = CE
	} else if order.CandyType == "AA" {
		price = AA
	} else if order.CandyType == "NT" {
		price = NT
	} else if order.CandyType == "DE" {
		price = DE
	} else if order.CandyType == "YR" {
		price = YR
	}
	change := order.Money - (price * order.CandyCount)
	return change

}

func processOrder(order Order) int {
	if order.Money <= 0 {
		return 0
	}
	if order.CandyCount <= 0 {
		return 0
	}
	if order.CandyType == "CE" {
		return 1
	} else if order.CandyType == "AA" {
		return 1
	} else if order.CandyType == "NT" {
		return 1
	} else if order.CandyType == "DE" {
		return 1
	} else if order.CandyType == "YR" {
		return 1
	}
	return 0

}

func returnResponse400(w http.ResponseWriter) {
	var response InlineResponse400

	// cs := C.CString("some error in input data!")
	// defer C.free(unsafe.Pointer(cs))
	// response.Error_ = C.GoString(C.ask_cow(cs))

	response.Error_ = "some error in input data!"
	convertByte, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		fmt.Println("Error MarshalIndent:", err)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, string(convertByte))
}

func returnResponse201(w http.ResponseWriter, change int32) {
	var response InlineResponse201

	cs := C.CString("Thank you!")
	defer C.free(unsafe.Pointer(cs))
	response.Thanks = C.GoString(C.ask_cow(cs))

	response.Change = change
	convertByte, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		fmt.Println("Error MarshalIndent:", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, string(convertByte))
}

func returnResponse402(w http.ResponseWriter, change int32) {
	var response InlineResponse402
	response.Error_ = "You need " + strconv.Itoa(int(-change)) + " more money!"

	// cs := C.CString("You need " + strconv.Itoa(int(-change)) + " more money!")
	// defer C.free(unsafe.Pointer(cs))
	// response.Error_ = C.GoString(C.ask_cow(cs))

	convertByte, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		fmt.Println("Error MarshalIndent:", err)
		return
	}
	w.WriteHeader(http.StatusPaymentRequired)
	fmt.Fprintf(w, string(convertByte))
}

func BuyCandy(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	resBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error ReadAll :", err)
		returnResponse400(w)
		return
	}
	var order Order
	err = json.Unmarshal(resBody, &order)
	if err != nil {
		fmt.Println("Error Unmarshal :", err)
		returnResponse400(w)
		return
	}

	if processOrder(order) == 0 {
		returnResponse400(w)
		return
	}
	// fmt.Printf("Money %d\n", order.Money)
	// fmt.Printf("CandyT %s\n", order.CandyType)
	// fmt.Printf("CandyCount %d\n", order.CandyCount)

	change := countChange(order)
	if change >= 0 {
		returnResponse201(w, change)

	} else if change < 0 {
		returnResponse402(w, change)
	}
}
