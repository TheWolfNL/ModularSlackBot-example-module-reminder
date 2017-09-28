package reminder

import (
	"github.com/thewolfnl/ModularSlackBot"
)

var reminders []reminder

// New function to return a new instance of this bot
func New() *bot.Module {
	reminder := bot.NewModule("ReminderModule", "0.0.2")

	// Define triggers
	reminder.AddTrigger("(?i)notify #[0-9]+.*", notify)
	reminder.AddTrigger("(?i)remind #[0-9]+.*", remind)

	return reminder
}

func notify(message *bot.Message) {
	if !message.IsBot() {
		reminders = append(reminders, reminder{1234, message.User})
		message.Respond("Setting reminder for '" + message.Text + "' @" + message.User)
		// fmt.Print(reminders)
	}
}

func remind(message *bot.Message) {
	if message.IsBot() {
		message.Respond("Notification from reminder")
	}
}
