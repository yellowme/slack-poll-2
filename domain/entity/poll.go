package entity

type PollMode int

type Poll struct {
	ID          string
	Question    string
	Owner       string
	Options     []string
	Mode        PollMode
	PollAnswers []PollAnswer
}

const (
	PollModeSingle   PollMode = 1
	PollModeMultiple PollMode = 2
)

func (pollMode PollMode) String() string {
	names := [...]string{
		"single",
		"multiple",
	}

	return names[pollMode]
}
