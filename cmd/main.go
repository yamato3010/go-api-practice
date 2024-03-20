package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    app.Get("/git", func(c *fiber.Ctx) error {
        resp, err := http.Get("https://api.github.com/users/yamato3010")
        if err != nil {
            return c.Status(500).SendString(err.Error())
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return c.Status(500).SendString(err.Error())
        }

        var result map[string]interface{}
        json.Unmarshal(body, &result)

        return c.JSON(result)
    })

    app.Get("/git/:username", func(c *fiber.Ctx) error {
        resp, err := http.Get("https://api.github.com/users/" + c.Params("username"))
        if err != nil {
            return c.Status(500).SendString(err.Error())
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return c.Status(500).SendString(err.Error())
        }

        var result map[string]interface{}
        json.Unmarshal(body, &result)

        return c.JSON(result)
    })

    app.Listen(":3003")
}
