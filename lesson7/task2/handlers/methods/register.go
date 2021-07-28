package methods

import (
	"encoding/json"
	"errors"
)

var Peoples = []string{"firstName", "SecondName"}

type registerParams struct {
	Name string `json:"name"`
}

func Register(params json.RawMessage) (interface{}, error) {
	p := registerParams{}
	if err := json.Unmarshal(params, &p); err != nil {
		return "", errors.New("wrong params sent, expected {name:string}")
	}
	Peoples = append(Peoples, p.Name)
	return "added", nil
}
