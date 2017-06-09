package lib

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const ENDPOINT = "https://factordb.com/api"

type FactorDB struct {
	Number int
	Result FactorDBResponse
}

func (f *FactorDB) Empty() bool {
	if f.Result.Status == "" {
		return true
	}
	return false
}

func (f *FactorDB) Connect() error {
	values := url.Values{}
	values.Add("query", fmt.Sprintf("%d", f.Number))
	resp, err := http.Get(fmt.Sprintf("%s?%s", ENDPOINT, values.Encode()))
	if err != nil {
		return errors.New("cannot connect" + ENDPOINT)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Empty Body")
	}

	response, err := ConvertToFactorDB(b)
	if err != nil {
		return errors.New("Cannot converting data")
	}

	f.Result = response
	return nil
}

func (f *FactorDB) GetId() (string, error) {
	if f.Empty() {
		return "", errors.New("Empty Result")
	}
	return f.Result.Id, nil
}

func (f *FactorDB) GetStatus() (string, error) {
	if f.Empty() {
		return "", errors.New("Empty Result")
	}
	return f.Result.Status, nil
}

func (f *FactorDB) GetFactorList() ([]int, error) {
	if f.Empty() {
		return []int{}, errors.New("Empty Result")
	}
	var sum int
	for _, f := range f.Result.Factors {
		sum += f.Power
	}

	var ret []int
	for _, f := range f.Result.Factors {
		for i := 0; i < f.Power; i++ {
			ret = append(ret, f.Number)
		}
	}
	return ret, nil
}
