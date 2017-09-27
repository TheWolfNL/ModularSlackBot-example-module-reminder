package reminder

import (
	"testing"

	"github.com/thewolfnl/ModularSlackBot"
)

func TestNew(t *testing.T) {
	module := New()

	if module.Name() != "ReminderModule" {
		t.Error("Expected module name ReminderModule, got ", module.Name())
	}

	if module.Version() != "0.0.1" {
		t.Error("Expected module version 0.0.1, got ", module.Version())
	}
}

func ExampleNotify() {
	module := New()
	module.HandleInput(bot.CreateMessage("notify #1"))
	// Output: Sending to #C2147483705
	// Response: Setting reminder for 'notify #1' @U2147483697
	// Message not sent to slack because slack api is not configured
}

func ExampleRemind() {
	module := New()
	message := bot.CreateMessage("remind #1")
	message.User = "USLACKBOT"
	module.HandleInput(message)
	// Output: Sending to #C2147483705
	// Response: Notification from reminder
	// Message not sent to slack because slack api is not configured
}
