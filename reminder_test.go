package reminder

import (
	"encoding/json"
	"testing"

	"github.com/thewolfnl/ModularSlackBot/bot"
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
	message, err := createMessage("notify #1")
	if err != nil {
		// exit
	}
	module.HandleInput(message)
	// Output: Sending to #C2147483705
	// Response: Setting reminder for 'notify #1' @U2147483697
	// Message not sent to slack because slack api is not configured
}

func ExampleRemind() {
	module := New()
	message, err := createMessage("remind #1")
	if err != nil {
		// exit
	}
	message.User = "USLACKBOT"
	module.HandleInput(message)
	// Output: Sending to #C2147483705
	// Response: Notification from reminder
	// Message not sent to slack because slack api is not configured
}

func createMessage(messageString string) (bot.Message, error) {
	message := bot.Message{}
	messageJson := `{
		"type": "message",
		"channel": "C2147483705",
		"user": "U2147483697",
		"text": "Hello world",
		"ts": "1355517523.000005"
		}`
	if err := json.Unmarshal([]byte(messageJson), &message); err != nil {
		return bot.Message{}, err
	}
	message.Text = messageString
	return message, nil
}

func TestCreateMessage(t *testing.T) {
	_, err := createMessage("Hello")
	if err != nil {
		t.Error("Could not create a test message")
	}
}
