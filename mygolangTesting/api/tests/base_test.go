package tests

import (
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"gotestinginteg/mygolangTesting/api/app"
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
