/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get random dad joke",
	Long: `This command fetch a random dad joke.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

type Joke struct{
	ID string `json:"id"`
	Joke string `json:joke`
	Status int `json:status`
}

func getRandomJoke(){
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)
	joke := Joke{}

	if err := json.Unmarshal(responseBytes,&joke); err!=nil{
		log.Println("Could not unMarshall joke responseBytes")
	}

	fmt.Println(joke.Joke)
}
func getJokeData(baseAPI string)[]byte{
	request,err:= http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
		)
	if err!=nil{
		log.Println("Failed struct get request")
	}

	request.Header.Add("Accept","application/json")
	request.Header.Add("User-Agent", "Dad Joke CLI (github.com/rweebs/dadjoke)")

	response,err:=http.DefaultClient.Do(request)

	if err !=nil{
		log.Println("Couldn't get request")
	}

	responseBytes,err:=ioutil.ReadAll(response.Body)

	if err != nil{
		log.Println("Couldn't read response body - %v",err)
	}
	return responseBytes
}
