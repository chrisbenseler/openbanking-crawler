package main

import (
	"fmt"
	"openbankingcrawler/application"
	"os"
)

func main() {

	mode := os.Getenv("MODE")

	if mode == "local" {
		fmt.Println("Start Open Banking Crawler application - local")

		application.NewLocal()
		return
	}

	fmt.Println("Start Open Banking Crawler application - web")

	application.NewWeb()
}
