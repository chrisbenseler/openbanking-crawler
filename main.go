package main

import (
	"fmt"
	"openbankingcrawler/application"
)

func main() {
	fmt.Println("Start Open Banking Crawler application - web")

	application.NewWeb()
}
