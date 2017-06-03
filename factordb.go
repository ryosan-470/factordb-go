package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/ryosan-470/factordb-go/handler"
)

const ENDPOINT = "https://factordb.com/api"

type FactorDB struct {
	Number int
	Result handler.FactorDbResponse
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
		return errors.New("empty body")
	}

	response := handler.ConvertToFactorDB(b)
	f.Result = response
	return nil
}

func main() {
	var factordb FactorDB

	factordb.Number = 20
	fmt.Printf("Before: %v\n", factordb)
	if err := factordb.Connect(); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("After:  %v\n", factordb)
}
