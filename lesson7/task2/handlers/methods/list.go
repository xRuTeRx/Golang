package methods

import (
	"encoding/json"
	"errors"
)

type listParams struct {
}

func List(params json.RawMessage) (interface{}, error) {
	p := listParams{}
	if err := json.Unmarshal(params, &p); err != nil {
		return "", errors.New("wrong params sent, expected {name:string}")
	}
	return Peoples, nil
}
