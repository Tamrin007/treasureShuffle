package main

import (
	"math/rand"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

// Result は Member のスライスを返す
type Result struct {
	Members []Member `json:"result"`
}

// Member は席番号と名前の組.
type Member struct {
	Num  int    `json:"num"`
	Name string `json:"name"`
}

// handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func shuffle(c echo.Context) error {
	result := new(Result)
	members := []string{
		"IanBrison",
		"RyotaKiuchi",
		"Tacha-S",
		"Tamrin007",
		"YuiOmata",
		"applepine1125",
		"atsumine",
		"dfree1645",
		"g-hyoga",
		"itezaP",
		"kazumasaSS",
		"kei-tamiya",
		"konnobu",
		"kozenumezawa",
		"sana3737",
		"simorgh3196",
		"skuwa229",
		"tacigar",
		"tajima12",
		"tantakan",
		"technicakidz",
		"tikasan",
		"ynarita",
	}

	n := len(members)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		members[i], members[j] = members[j], members[i]
	}

	for i, member := range members {
		var m Member
		m.Num = i
		m.Name = member
		result.Members = append(result.Members, m)
	}
	return c.JSON(http.StatusOK, result)
}

func main() {
	// Echo instance
	e := echo.New()

	// Route
	e.GET("/", hello)

	e.GET("/shuffle", shuffle)

	// Start server
	e.Run(standard.New(":" + os.Getenv("PORT")))
}
