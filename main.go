package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tutorial-embed-react-with-go/web"

	echo "github.com/labstack/echo/v4"
)

type carro struct {
	Modelo modelo
	Marca  string
	Status bool
}

type modelo struct {
	Nome string
}

func (c *carro) partida() {

	c.Status = true
}

func main() {

	var veiculo = carro{

		Modelo: modelo{Nome: "gol"},
		Marca:  "vw",
		Status: false,
	}

	veiculo.partida()

	app := echo.New()

	web.RegisterHandlers(app)
	app.GET("/api", func(c echo.Context) error {

		veiculoJson, err := json.Marshal(veiculo)

		if err != nil {
			fmt.Println(veiculoJson)
			return c.String(http.StatusInternalServerError, "erro")

		} else {

			return c.JSON(http.StatusOK, veiculo)

		}

	})

	app.Logger.Fatal(app.Start(":8080"))

}
