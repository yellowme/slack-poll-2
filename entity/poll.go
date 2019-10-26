package entity

type PollMode string

const (
	Single   PollMode = "single"
	Multiple PollMode = "multiple"
)

type Poll struct {
	ID       string
	Question string
	Owner    string
	Channel  string
	Options  []string
	Mode     PollMode
}
