package main

import (
	"log"
	"super-gen/internal/generator"
)

func main() {
	log.Println("run main")
	generator := generator.New("generated-go-files")
	generator.Run()
}
