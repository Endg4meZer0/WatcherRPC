package wrpcTypes

type ProcessToWatch struct {
	ProcessName    string
	Details        string
	State          string
	LargeImageKey  string
	LargeImageText string
	SmallImageKey  string
	SmallImageText string
	UseTimestamp   bool
	Buttons        []ActivityButtons
}

type ActivityButtons struct {
	Text string
	Url  string
}
