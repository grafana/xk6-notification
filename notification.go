// Package notification enables messages to be sent to a Slack workspace via webhook
package notification

import (
	"github.com/containrrr/shoutrrr"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/notification", new(NOTIFICATION))
}

// NOTIFICATION is the k6 notification extension.
type NOTIFICATION struct{}

// Send a notification using Shoutrr
func (*NOTIFICATION) Send(url string, msg string) {
	// TODO: Format? Remove? Shoutrr logs
	err := shoutrrr.Send(url, msg)
	if err != nil {
		ReportError(err, "Failed to send notification")
	}
}
