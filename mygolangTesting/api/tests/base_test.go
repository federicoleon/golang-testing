package tests

import (
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/golang-testing/mygolangTesting/api/app"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	fmt.Println("now starting application")
	go app.StartApp()
	os.Exit(m.Run())
	fmt.Println("Application STarted")
}
