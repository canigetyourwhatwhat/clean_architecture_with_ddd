package main

import (
	"clean_architecture_with_ddd/cmd"
	"log"
)

func main() {
	e := cmd.Run()
	log.Fatal(e.Start(":9000"))
}
