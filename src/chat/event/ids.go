package event

const (
	ServerCreate Id = iota
	ServerStartListen
	ServerStartRun
	ClientOnline
	ClientOffline
	ServerEnd
)
