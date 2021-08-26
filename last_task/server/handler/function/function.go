package function

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func WeatherFromOpenweathermap(w http.ResponseWriter, cityName string) ([]byte, error) {
	url := fmt.Sprintf(os.Getenv("WEATHER_HOST"), cityName, os.Getenv("ID_KEY"))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	var res []byte
	if resp.StatusCode != 200 {
		message := result["message"]
		res, err = json.Marshal(message)
		if err != nil {
			return nil, err
		}
		http.Error(w, string(res), resp.StatusCode)

	} else {
		res, err = json.Marshal(result["main"])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(res, &result)
		if err != nil {
			return nil, err
		}
		res, err = json.Marshal(result["temp"])
		if err != nil {
			return nil, err
		}
		w.Write(res)
	}
	return res, nil
}
