package observer

type Observer interface {
	UpdateOnEvent(string)
	GetId() string
}
