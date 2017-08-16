package main

import (
	"encoding/json"
)

/*
data = {
	"status": "ok",
	"code": 200,
	"body": "herp"
}
*/

type Response struct {
	Status string `json:"status"`
	Code   string `json:"code"`
	Body   string `json:"data"`
}

func main() {
	res := &Response{}
	json.Unmarshal(data, res)

	// Now you can interact with the data as a struct
}
