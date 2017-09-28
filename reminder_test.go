package reminder

import (
	"testing"

	"github.com/thewolfnl/ModularSlackBot"
	"github.com/thewolfnl/ModularSlackBot/testLib"
)

func setup() (*botTestLib.MockBot, *bot.Module) {
	module := New()
	mock := botTestLib.NewMockBot()
	mock.AddModule(module)
	return mock, module
}

func TestNew(t *testing.T) {
	_, module := setup()

	if module.Name() != "ReminderModule" {
		t.Error("Expected module name ReminderModule, got ", module.Name())
	}

	if module.Version() != "0.0.2" {
		t.Error("Expected module version 0.0.2, got ", module.Version())
	}
}

func ExampleNotify() {
	mock, module := setup()
	module.HandleInput(mock.CreateMessage("notify #1"))
	// Output: Sending to #C2147483705
	// Response: Setting reminder for 'notify #1' @U2147483697
	// Message not sent to slack because slack api is not configured
}

func ExampleRemind() {
	mock, module := setup()
	message := mock.MockMessage(mock.CreateMessageEvent("remind #1"), botTestLib.ReturnTrue, botTestLib.ReturnFalse, botTestLib.ReturnFalse)
	module.HandleInput(message)
	// Output: Sending to #C2147483705
	// Response: Notification from reminder
	// Message not sent to slack because slack api is not configured
}
