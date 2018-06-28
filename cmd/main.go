package main

import (
	"fmt"

	"github.com/na1s/mfa/internal/pkg/config"
	"github.com/na1s/mfa/internal/pkg/data/events"
)

func main() {
	config := &config.DummyConfiguration{}
	events := events.GetEvents(config)
	fmt.Print(events)
}
