package main

import (
	"futurisme-api/cmd/commands"
)

func main() {
	// Memanggil inisiator Cobra CLI
	commands.Execute()
}

// go run main.go start --dev
// go run main.go seed
// go run main.go migrate
