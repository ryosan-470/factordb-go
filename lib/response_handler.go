package lib

import (
	"encoding/json"
	"errors"
	"strconv"
)

type FactorDBResponse struct {
	Status  string
	Id      string
	Factors []Factor
}

type Factor struct {
	Number int
	Power  int
}

func ConvertToFactorDB(b []byte) (FactorDBResponse, error) {
	var base interface{}
	err := json.Unmarshal(b, &base)
	if err != nil {
		return FactorDBResponse{}, errors.New("Cannot parse the input")
	}
	s := base.(map[string]interface{})

	var factor FactorDBResponse
	factor.Status = s["status"].(string)
	factor.Id = s["id"].(string)

	factors := s["factors"].([]interface{})

	for _, f := range factors {
		tmp := f.([]interface{})
		number, _ := strconv.Atoi(tmp[0].(string))
		power := int(tmp[1].(float64))
		factor.Factors = append(factor.Factors, Factor{number, power})
	}

	return factor, nil
}
