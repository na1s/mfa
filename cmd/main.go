package main

import (
	"fmt"

	"github.com/na1s/mfa/internal/pkg/data/events"
)

func main() {
	events := events.Connect()
	fmt.Println(events)
}
