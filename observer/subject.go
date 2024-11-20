package observer

type Subject interface {
	RegisterObserver(o Observer)
	UnregisterObserver(o Observer)
	NotifyObservers()
}
