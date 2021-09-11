package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Data struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	var data Data
	jData, _ := json.Marshal(data)
	zLogger := zerolog.New(os.Stdout)
	zLogger.Debug().
		Interface("data", data).
		Bools("bools", []bool{true, false, true, true}).
		Bytes("data_marshal_bytes", jData).
		Interface("data_marshal", jData).
		Msg("test message")
}
