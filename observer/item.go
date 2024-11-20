package observer

import "fmt"

type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func (i *Item) RegisterObserver(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) UnregisterObserver(o Observer) {
	for index, observer := range i.observerList {
		if observer.GetId() == o.GetId() {
			i.observerList = append(i.observerList[:index], i.observerList[index+1:]...)
			break
		}
	}
}

func (i *Item) NotifyObservers() {
	for _, observer := range i.observerList {
		observer.UpdateOnEvent(i.name)
	}
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) UpdateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.NotifyObservers()
}
