package webserver

type GinServerMode int

const (
	DebugMode GinServerMode = iota
	ReleaseMode
	TestMode
)
