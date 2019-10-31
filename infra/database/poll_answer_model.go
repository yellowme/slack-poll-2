package database

type PollAnswerModel struct {
	ID     string
	Option string
	Owner  string
	Poll   PollModel
	PollID string
}
