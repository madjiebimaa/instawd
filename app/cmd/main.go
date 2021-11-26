package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/madjiebimaa/go-random-quotes/helper"
	"github.com/madjiebimaa/go-random-quotes/model/web"
)

func main() {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getRandom := getCmd.Bool("random", false, "Get random quote")
	getAll := getCmd.Bool("all", false, "Get all quotes")
	getID := getCmd.String("id", "", "Quote ID")

	if len(os.Args) < 2 {
		fmt.Println("expected 2 subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, getRandom, getAll, getID)
	default:
	}
}

func HandleGet(getCmd *flag.FlagSet, random *bool, all *bool, id *string) {

	getCmd.Parse(os.Args[2:])

	if !*random && *all && *id == "" {
		fmt.Print("id is required or specify --all for all quotes or --random for random quote")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *random {
		res, err := http.Get("http://localhost:3000/api/quotes/random-quote")
		helper.PanicIfError(err)
		defer res.Body.Close()

		var webResponse web.WebResponse
		err = json.NewDecoder(res.Body).Decode(&webResponse)
		helper.PanicIfError(err)

		quote := (webResponse.Data).(map[string]interface{})

		fmt.Printf("\"%s\" ~ %s\n", quote["content"], quote["author"])
		return
	}

	if *all {

		return
	}

	if *id != "" {

		return
	}

}
