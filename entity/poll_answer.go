package entity

type PollAnswer struct {
	Option string
	Owner  string
	Poll   Poll
}
