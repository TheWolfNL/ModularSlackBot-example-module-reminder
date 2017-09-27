package reminder

import (
	"github.com/thewolfnl/ModularSlackBot"
)

var reminders []reminder

// New function to return a new instance of this bot
func New() ReminderModule {
	reminder := ReminderModule{bot.NewModule("ReminderModule", "0.0.1")}

	// Define triggers
	reminder.AddTrigger("(?i)notify #[0-9]+.*", reminder.notify)
	reminder.AddTrigger("(?i)remind #[0-9]+.*", reminder.remind)

	return reminder
}

func (module *ReminderModule) notify(message *bot.Message) error {
	if !message.IsBot() {
		reminders = append(reminders, reminder{1234, message.User})
		module.Respond("Setting reminder for '" + message.Text + "' @" + message.User)
		// fmt.Print(reminders)
	}
	return nil
}

func (module *ReminderModule) remind(message *bot.Message) error {
	if message.IsBot() {
		module.SetChannel(message.Channel)
		module.Respond("Notification from reminder")
	}
	return nil
}
