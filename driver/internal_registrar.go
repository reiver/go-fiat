package fiatdriver

import (
	"sync"
)

type internalRegistrar struct {
	m map[string]Driver
	mutex sync.RWMutex
}

func (receiver *internalRegistrar) Register(name string, driver Driver) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == driver {
		return errNilDriver
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.m {
		receiver.m = map[string]Driver{}
	}

	if _, ok := receiver.m[name]; ok {
		return errFound
	}

	receiver.m[name] = driver

	return nil
}

func (receiver *internalRegistrar) Obtain(name string) (Driver, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	if nil == receiver.m {
		receiver.m = map[string]Driver{}
	}

	driver, ok := receiver.m[name]
	if !ok {
		return nil, internalNotFoundComplainer{name}
	}
	if nil == driver {
		return nil, errNilDriver
	}

	return driver, nil
}
