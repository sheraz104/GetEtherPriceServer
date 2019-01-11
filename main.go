package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/etherPrice", func(c *gin.Context) {

		price := make(chan string)
		go getPrice(price)

		c.JSON(http.StatusOK, gin.H{
			"ETH-USD": "$ " + <-price,
		})
	})

	router.Run(":8888")
}

func getPrice(price chan string) {
	response, err := http.Get("https://api.etherscan.io/api?module=stats&action=ethprice&apikey=YourApiKeyToken")
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	var result map[string]interface{}

	json.Unmarshal([]byte(contents), &result)

	price <- result["result"].(map[string]interface{})["ethusd"].(string)

}
