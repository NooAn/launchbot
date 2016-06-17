package main

import (
  "fmt"
  "os/exec"
  "strings"
  "github.com/telegram-bot-api"
)


// only for linux server
func main() {
  var s string 
  bot, _ := tgbotapi.NewBotAPI(TOKEN_LAUNCH)
  bot.Debug = false
  fmt.Println("Authorized on account %s", bot.Self.UserName)
  var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
  ucfg.Timeout = 160
  updates, _ := bot.GetUpdatesChan(ucfg)
  for update := range updates {
    UserName := update.Message.From.UserName
    ChatID := update.Message.Chat.ID
    Text := update.Message.Text
    s ="I am server... and I say "
    if strings.ContainsAny(NAME, UserName) && Text == PASS {
      _, err := exec.Command(LAUNCH_COMMAND).Output()
      if err != nil {
        s += fmt.Sprintf("%s", err)
      } else {
        s += " bot launched"
      }
    }
    msg := tgbotapi.NewMessage(ChatID, s)
    bot.Send(msg)
  }
}