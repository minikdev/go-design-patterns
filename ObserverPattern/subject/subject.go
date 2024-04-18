package subject

import "example.com/observer-pattern/observer"

type ISubject interface {
	RegisterObserver(o observer.IObserver)
	RemoveObserver(o observer.IObserver)
	NotifyObservers()
}
