package event

const (
	ServerCreate Id = iota
	ServerStartListen
	ServerStartRun
	ClientOnline
	ClientOnMessage
	ClientOffline
	ServerEnd
)
