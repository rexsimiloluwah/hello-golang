package client

import (
	"os"
	"log"
	"net/http"
	"encoding/json"
	"../models"
)


func FetchCryptoResponse(fiat string, crypto string) (string,error){
	API_KEY := os.Getenv("NOMIC_API_KEY")
	BASE_URL := "https://api.nomics.com/v1/currencies/ticker"
	ENDPOINT:= BASE_URL+"?key="+API_KEY+"&interval=1d&ids="+crypto+"&convert="+fiat
	resp,err:= http.Get(ENDPOINT)
	if err != nil{
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var apiResponse models.ApiResponse 
	decodeErr := json.NewDecoder(resp.Body).Decode(&apiResponse)
	if decodeErr != nil{
		log.Fatalln(decodeErr)
	}

	return apiResponse.ResponseTable(), nil
}