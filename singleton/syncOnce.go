package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type singleOnce struct {
	resoureInitialized bool
}

var singleOnceInstance *singleOnce

func GetSingleInstance() *singleOnce {
	if singleOnceInstance == nil {
		once.Do(func() {
			fmt.Println("initializing the resource")
			singleOnceInstance = &singleOnce{
				resoureInitialized: true,
			}
		})
	} else {
		fmt.Println("resource already initialized")
	}
	return singleOnceInstance
}
