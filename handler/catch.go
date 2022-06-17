package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Option struct {
	Name string `json:"name" xml:"name" form:"name"`
}

func Catch(c *fiber.Ctx) error {
	var result = rand.Intn(2)
	var option Option
	var status bool
	if err := c.BodyParser(&option); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

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
		return c.Status(400).JSON(fiber.Map{
			"Message": err.Error(),
			"status":  false,
		})
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Message": err.Error(),
			"status":  false,
		})
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Message": err.Error(),
			"status":  false,
		})
	}

	m := map[string]interface{}{}
	if err := json.Unmarshal(body, &m); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Message": err.Error(),
			"status":  false,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"Message": "done",
		"data":    m,
		"status":  status,
	})
}
