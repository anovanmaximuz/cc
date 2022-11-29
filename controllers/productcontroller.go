package controllers

import (
	"crypto-community/database"
	"crypto-community/entities"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Status  string
	Message string
	Data    interface{}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Create(&product)
	json.NewEncoder(w).Encode(product)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if checkIfProductExists(productId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entities.Product
	database.Instance.First(&product, productId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var treaders []entities.Community
	//database.Instance.Find(&treaders)
	database.Instance.Raw("SELECT uid, display_name as name,image as photo,win,roe,day,premium_type as premium,monetize,kyc_verified as verified FROM user_data WHERE use_app='yes' AND sso<>'device' LIMIT 0,20").Scan(&treaders)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res2 := Response{
		Status: "OK",
		Message: "no error",
		Data: treaders
	}
	json.NewEncoder(w).Encode(res2)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if checkIfProductExists(productId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entities.Product
	database.Instance.First(&product, productId)
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Save(&product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	productId := mux.Vars(r)["id"]
	if checkIfProductExists(productId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entities.Product
	database.Instance.Delete(&product, productId)
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
}

func checkIfProductExists(productId string) bool {
	var product entities.Product
	database.Instance.First(&product, productId)
	if product.ID == 0 {
		return false
	}
	return true
}
