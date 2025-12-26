package core

type AppState struct {
	DeviceConnected bool
	CurrentPage     string
	Wheel           *Wheel
}

func NewAppState() *AppState {
	return &AppState{
		DeviceConnected: false,
		CurrentPage:     "",
		Wheel:           nil,
	}
}
