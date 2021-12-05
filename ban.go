package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

func ban(m *tb.Message) {
	if m.Private() {
		b.Reply(m, "This command is for groups.")
		return
	}
	user, _ := get_user(m)
	if user == nil {
		b.Reply(m, "You dont seem to be referring to a user or the ID specified is incorrect..")
	}
	err := b.Ban(m.Chat, &tb.ChatMember{
		User: user,
	})
	if err == nil {
		b.Reply(m, "<b>"+user.FirstName+"</b> was banned. ~")
		return
	}
	b.Reply(m, "Failed to ban, "+string(err.Error()))
}
