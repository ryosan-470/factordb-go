package handler

import (
	"encoding/json"
	"log"
	"strconv"
)

type FactorDB struct {
	Status  string
	Id      string
	Factors []Factor
}

type Factor struct {
	Number int
	Power  int
}

func ConvertToFactorDB(b []byte) FactorDB {
	var base interface{}
	err := json.Unmarshal(b, &base)
	if err != nil {
		log.Fatal("Cannot parse your input")
	}
	s := base.(map[string]interface{})

	var factor FactorDB
	factor.Status = s["status"].(string)
	factor.Id = s["id"].(string)

	factors := s["factors"].([]interface{})

	for _, f := range factors {
		tmp := f.([]interface{})
		number, _ := strconv.Atoi(tmp[0].(string))
		power := int(tmp[1].(float64))
		factor.Factors = append(factor.Factors, Factor{number, power})
	}

	return factor
}
