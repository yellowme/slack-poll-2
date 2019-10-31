package database

type PollModel struct {
	ID       string
	Question string
	Owner    string
	Options  string
	Mode     string
}
