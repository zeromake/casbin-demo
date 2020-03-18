package main

import (
	"log"
	"net/http"

	"github.com/casbin/casbin/v2"
)

func main() {
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		log.Fatal(err)
	}
	e.EnableLog(true)
	ok, err := e.Enforce("admin", "/v1/admin/manager/1", http.MethodPut)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ok)
}
