package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tomsteele/boom"
)

type vError struct {
	Source string   `json:"source"`
	Keys   []string `json:"keys"`
}

func main() {
	validationError := map[string]interface{}{"validation": vError{
		Source: "payload",
		Keys:   []string{"email"},
	}}
	err := boom.BadRequest("invalid payload", validationError)
	fmt.Println(err)
	data, err := json.Marshal(err)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	err = boom.BadImplementation(errors.New("sql: no rows found in result set"))
	fmt.Println(err)
	data, err = json.Marshal(err)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

}
