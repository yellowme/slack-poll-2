package webserver

import (
	"strings"

	"github.com/jerolan/slack-poll/core/entity"
)

func ConvertCommandToStruct(command string) (poll entity.Poll) {
	sanitizedCommand := removeDoubleCuotes(command)
	poll.Mode = entity.PollModeSingle

	if strings.Contains(sanitizedCommand, "-m") {
		poll.Mode = entity.PollModeMultiple
	}

	var options []string
	dirtyOptions := strings.Split(strings.ReplaceAll(sanitizedCommand, "-m", ""), `"`)

	for _, option := range dirtyOptions {
		pseudoOption := strings.Trim(option, " ")
		if pseudoOption != "" {
			options = append(options, pseudoOption)
		}
	}

	poll.Options = options[2:]
	poll.Question = options[1]
	return poll
}

func removeDoubleCuotes(str string) string {
	clean := strings.ReplaceAll(str, "\u201D", `"`)
	clean = strings.ReplaceAll(clean, "\u201C", `"`)
	return clean
}
