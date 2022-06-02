package excuses

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func Bahana(c echo.Context) (err error) {
	resp, err := http.Get("https://techy-api.vercel.app/api/text")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Println(sb)
	return c.JSON(http.StatusOK, sb)
}
