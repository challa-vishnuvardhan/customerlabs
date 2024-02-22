package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Worker(w http.ResponseWriter, r *http.Request) {

	input := make(chan string)
	go worker(input)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	input <- string(body)
	defer close(input)
	w.Header().Add("content-type", "application/json")
	w.Write([]byte(<-input))

}

func worker(input chan string) {
	for inputBody := range input {
		response := Response{}
		var inputMap map[string]string
		err := json.Unmarshal([]byte(inputBody), &inputMap)
		if err != nil {
			log.Fatal(err)
		}
		i := 1
		attKey := "atrk"
		attributes := make(map[string]TypeValue)
		for {
			if key, ok := inputMap[fmt.Sprintf("%v%v", attKey, i)]; ok {
				attributes[key] = TypeValue{
					Value: inputMap[fmt.Sprintf("atrv%v", i)],
					Type:  inputMap[fmt.Sprintf("atrt%v", i)],
				}

			} else {
				break
			}
			i++
		}

		j := 1
		uatrkKey := "uatrk"
		traits := make(map[string]TypeValue)
		for {
			if key, ok := inputMap[fmt.Sprintf("%v%v", uatrkKey, j)]; ok {
				traits[key] = TypeValue{
					Value: inputMap[fmt.Sprintf("uatrv%v", j)],
					Type:  inputMap[fmt.Sprintf("uatrt%v", j)],
				}

			} else {
				break
			}
			j++
		}

		response.Event = inputMap["ev"]
		response.EventType = inputMap["et"]
		response.AppId = inputMap["id"]
		response.UserId = inputMap["uid"]
		response.MessageId = inputMap["mid"]
		response.PageTitle = inputMap["t"]
		response.PageUrl = inputMap["p"]
		response.BrowserLanguage = inputMap["l"]
		response.ScreenSize = inputMap["sc"]
		response.Attributes = attributes
		response.Traits = traits
		result, err := json.Marshal(response)

		if err != nil {
			log.Fatal(err)
		}
		input <- string(result)

		_, err = http.Post("https://webhook.site/617821ee-5d69-42fa-b6ef-744bf8b34653", "application/json",
			bytes.NewBuffer(result))
		if err != nil {
			log.Fatal(err)
		}
	}
}
