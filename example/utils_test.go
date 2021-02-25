package example_test

import (
	"fmt"

	"github.com/im-jinsu/easy-logging-level/utils/discord"
)

func TestDiscord() {
	s := new(discord.MsgBlock)
	s.Content = fmt.Sprintf("테스트입니다.")
	s.Username = "Test Agent"

	if err := s.Send(); err != nil {
		panic(err)
	}
}
