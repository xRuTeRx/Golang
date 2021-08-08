package methods

import (
	"encoding/json"
	"errors"
)

const errWromgParam = "wrong params sent, expected {name:string}"

var Peoples = []string{"firstName", "SecondName"}

type registerParams struct {
	Name string `json:"name"`
}

func Register(params json.RawMessage) (interface{}, error) {
	p := registerParams{}
	if err := json.Unmarshal(params, &p); err != nil {
		return "", errors.New(errWromgParam)
	}
	Peoples = append(Peoples, p.Name)
	return "added", nil
}
