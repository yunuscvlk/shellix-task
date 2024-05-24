package utils

import (
	"encoding/json"
	"log"

	"github.com/yunuscvlk/shellix-task/internal/models/result"
)

func Result(status bool, content string) []byte {
	r := new(result.Result)

	r.Status = status
	r.Content = content

	data, err := json.Marshal(r)

    if err != nil {
        log.Println(err)
    }

	return data
}