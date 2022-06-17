package usecase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"mini-project-pikachu/models"
	"net/http"
	"os"
)

func Catch(option models.Option) (map[string]interface{}, bool, error) {
	var result = rand.Intn(2)
	var status bool
	m := map[string]interface{}{}

	if result == 1 {
		fmt.Println("you success catch pokemon")
		status = true
	} else {
		fmt.Println("you failed catch pokemon")
		status = false
	}

	URI := os.Getenv("POKEAPI_URL") + "/api/v2/pokemon/" + option.Name
	req, err := http.NewRequest("GET", URI, nil)
	if err != nil {
		return m, status, err

	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return m, status, err

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return m, status, err

	}

	if err := json.Unmarshal(body, &m); err != nil {
		return m, status, err

	}
	return m, status, err

}
