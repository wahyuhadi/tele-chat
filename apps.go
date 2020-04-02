package main

import (
	"TGetupdates/model"
	"encoding/json"
	"fmt"
	"os"

	httpreq "github.com/wahyuhadi/httpreq"
)

var (
	URI  = "https://api.telegram.org/bot"
	FUNC = "/getUpdates"
)

func main() {
	// export BOTKEY=xxxxxxxxxx
	Botkey := os.Getenv("BOTKEY")
	if Botkey == "" {
		fmt.Println("Botkey not found", Botkey)
		os.Exit(1)
	}

	URI = URI + Botkey + FUNC
	r, err := httpreq.Req(URI)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer r.Body.Close()
	Res := new(model.GetUpdatesResponse)
	json.NewDecoder(r.Body).Decode(Res) // Decode for Interface Response
	ParsingResponse(Res)
}

func ParsingResponse(Res *model.GetUpdatesResponse) {
	for _, data := range Res.Result {
		if data.Message.Text != "" {
			user := data.Message.From.Username
			fmt.Println("----- > ["+user+"] ->", data.Message.Text)
			if data.Message.Document.FileID != "" {
				fmt.Println("[FOUND DOCUMENT]", data.Message.Document.FileName)
			}
		}

	}
}
