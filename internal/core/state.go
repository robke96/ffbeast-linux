package core

type AppState struct {
	DeviceConnected bool
	CurrentPage     string
}

func NewAppState() *AppState {
	return &AppState{
		DeviceConnected: false,
		CurrentPage:     "",
	}
}
