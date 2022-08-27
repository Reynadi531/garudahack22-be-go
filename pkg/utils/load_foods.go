package utils

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io/ioutil"
)

var Foods []string

func LoadFoods() {
	file, err := ioutil.ReadFile("foods.json")
	if err != nil {
		log.Error().Err(err).Msg("Error when reading foods.json")
	}
	if err := json.Unmarshal(file, &Foods); err != nil {
		log.Error().Err(err).Msg("Error when unmarshalling foods.json")
	}
}
