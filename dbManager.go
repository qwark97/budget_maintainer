package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mockedDBMutex = sync.Mutex{}
	mockedDB      = make(map[int]Operation)
	lastMockedID  = 0
)

func saveOperation(data Operation) error {
	mockedDBMutex.Lock()
	defer mockedDBMutex.Unlock()
	lastMockedID++
	id := lastMockedID
	data.ID = id
	data.Timestamp = int(time.Now().Unix())
	mockedDB[id] = data
	return nil
}

func eraseOperation(id int) error {
	if _, ok := mockedDB[id]; ok {
		mockedDBMutex.Lock()
		defer mockedDBMutex.Unlock()
		delete(mockedDB, id)
		return nil
	} else {
		return fmt.Errorf("no such id")
	}

}

func loadAllOperations() (Operations, error) {
	operations := Operations{}
	for _, o := range mockedDB {
		operations = append(operations, o)
	}
	return operations, nil
}

func alterAssets(amount int) error {
	return nil
}

func loadAssets() (Assets, error) {
	return Assets{}, nil
}
