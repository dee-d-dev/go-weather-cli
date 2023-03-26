package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	checkCmd = &cobra.Command{
		Use:   "check",
		Short: "Check weather info",
		Long:  `Check waeather info`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				for _, value := range args {
					fmt.Println()
					showWeatherData(value)
				}
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(checkCmd)
}

type WeatherInfo struct {
	Location struct {
		Name    string `json:"name"`
		Region  string `json:"region"`
		Country string `json:"country"`
	}
	Current struct {
		Temp_c float64 `json:"temp_c"`
		Temp_f float64 `json:"temp_f"`
	}
}

func showWeatherData(value string) {
	query := value

	url := "http://api.weatherapi.com/v1/current.json?q=" + query

	// var data map[string]interface{}

	data := &WeatherInfo{}
	responseByte := getWeatherData(url)
	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		fmt.Println("Unable to unmarshal response")
	}
	temperature := int64(data.Current.Temp_c)
	name := data.Location.Name
	message := fmt.Sprintf("the temperature in %s is %d°c", name, temperature)
	fmt.Println(message)

}

func getWeatherData(url string) []byte {
	client := &http.Client{}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	API_KEY := os.Getenv("API_KEY")

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Unable to get weather data")
	}

	req.Header.Add("Key", API_KEY)

	response, err := client.Do(req)
	if err != nil {

		fmt.Println(err)
	}

	defer response.Body.Close()
	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Print the response body
	return body
}
