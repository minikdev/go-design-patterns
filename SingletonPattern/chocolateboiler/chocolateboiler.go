package chocolateboiler

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}
var instance *ChocolateBoiler

type ChocolateBoiler struct {
	empty  bool
	boiled bool
}

func GetInstance() *ChocolateBoiler {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Creating a new instance now.")
			instance = &ChocolateBoiler{
				empty:  true,
				boiled: false,
			}
		}
	} else {
		fmt.Println("Returning existing instance.")
	}
	return instance
}

func (cb *ChocolateBoiler) Boil() {
	if !cb.empty && !cb.boiled {
		fmt.Println("Boiling chocolate and milk.")
		cb.boiled = true
	}
}
func (cb *ChocolateBoiler) Drain() {
	if !cb.empty && cb.boiled {
		fmt.Println("Draining boiled milk and chocolate.")
		cb.empty = true
	}
}
func (cb *ChocolateBoiler) Fill() {
	if cb.empty {
		fmt.Println("Filling the boiler with a milk/chocolate mixture.")
		cb.empty = false
	}
}
