/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("random called")
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := fetchJoke(url)
	joke := &Joke{}
	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		log.Printf("Error unmarshaling JSON: %v", err)
	}
	fmt.Println(string(joke.Joke))

}

func fetchJoke(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		log.Printf("Error creating HTTP request: %v", err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "dadjoke-cli (github.com/DiliniMW/dadjoke)")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error making HTTP request: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code: %d", response.StatusCode)
	}
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
	}
	return responseBytes

}
