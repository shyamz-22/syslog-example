package main

import (
	"fmt"
	"gopkg.in/mcuadros/go-syslog.v2"
	"gopkg.in/mcuadros/go-syslog.v2/format"
	"log"
	"strings"
)

func main() {
	// Get Creative here, make the patterns configurable via some
	// persistent storage
	patterns := []string{"Error", "ConnectionRefused"}
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.Automatic)
	server.SetHandler(handler)
	addr := "0.0.0.0:5514"
	if err := server.ListenTCP(addr); err != nil {
		log.Fatalf("Unable to resolve addr %s: %v", addr, err)
	}
	if err := server.Boot(); err != nil {
		log.Fatalf("cannot start syslog server, %v", err)
	}

	go alert(patterns, channel)

	server.Wait()
}


func alert(patterns []string, channel syslog.LogPartsChannel) {
	for logParts := range channel {
		fmt.Printf("📨 received %v\n", logParts)
		for _, pattern := range patterns {
			if matched := match(pattern, logParts["content"].(string)); matched {
				sendAlert(logParts)
			}
		}
	}
}

// Matcher can match on regular expression
func match(pattern string, message string) bool {
	return strings.Contains(message, pattern)
}

// Alerter could be some real thing
func sendAlert(parts format.LogParts) {
	fmt.Printf("🚨Alert Sent🚨 for %v\n", parts["content"])
}
