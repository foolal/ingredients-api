package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Product struct {
	Product struct {
		IngredientsText string `json:"ingredients_text"`
	} `json:"product"`
}

func openFoodFacts() {
	url := "https://world.openfoodfacts.org/api/v2/search?product=lays&fields=code,product_name,ingredients"

	fmt.Println(url)
	// Send GET request to the OpenFoodFacts API
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// Parse the JSON response
	var data struct {
		Products []Product `json:"products"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	fmt.Println(string(body))

	// // Collect unique ingredients
	// ingredients := make(map[string]bool)
	// for _, product := range data.Products {
	// 	fmt.Println(product)
	// 	ingredients[product.Product.IngredientsText] = true
	// }

	// // Print unique ingredients
	// fmt.Println("Unique Ingredients:")
	// for ingredient := range ingredients {
	// 	fmt.Println(ingredient)
	// }
}

func veganCheck() {

	var hello []byte
	bodyReader := bytes.NewReader(hello)

	prefix := "https://api.vegancheck.me/v0/product/"
	product := "3274080005003"
	url := prefix + product

	fmt.Println(url)

	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		fmt.Printf("Error setting new client: %v\n", err)
		return
	}

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return
	}
	defer res.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	fmt.Println(string(body))
}

func main() {
	veganCheck()
}
