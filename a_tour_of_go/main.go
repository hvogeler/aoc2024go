package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	keepDataCmdChan := make(chan Command, 100)
	gatherDataCmdChan := make(chan Command, 100)
	dataChan := make(chan *StateData, 100)

	waitGroup.Add(1)
	go KeepData(keepDataCmdChan, dataChan, &waitGroup)
	waitGroup.Add(1)
	go GatherData(gatherDataCmdChan, keepDataCmdChan, dataChan, &waitGroup)
	time.Sleep(time.Duration(3 * time.Second))

	fmt.Printf("Number of Go Routines: %d\n", runtime.NumGoroutine())
	fmt.Printf("Number of Threads: %d\n", runtime.GOMAXPROCS(0))
	iterations := 40
	if len(os.Args) == 2 {
		iterations, _ = strconv.Atoi(os.Args[1])
	}
	for i := 0; i < iterations; i++ {
		keepDataCmdChan <- Read
		data := <-dataChan
		slog.Info("main - Received Data: ", "Data", data)
		time.Sleep(time.Duration(4 * time.Second))
	}

	gatherDataCmdChan <- Quit
	keepDataCmdChan <- Quit
	waitGroup.Wait()

}

type Command int

const (
	Quit Command = iota
	Read
	Update
)

type State int

const (
	Unknown State = iota
	Heating
	Ready
)

func (state State) String() string {
	switch state {
	case State(Heating):
		return "Heating"
	case State(Ready):
		return "Ready"
	default:
		return "Unknown"
	}
}

type StateData struct {
	headTemp  int
	waterTemp int
	state     State
}

func (stateData StateData) String() string {
	return fmt.Sprintf("Temp: Head(%d) Water(%d), State: %v", stateData.headTemp, stateData.waterTemp, stateData.state)
}

func KeepData(keepDataCmdChan <-chan Command, dataChan chan *StateData, waitGroup *sync.WaitGroup) {
	slog.Info("KeepData - Starting KeepData")
	defer waitGroup.Done()
	state := new(StateData)
Loop:
	for i := 0; ; i++ {
		cmd := <-keepDataCmdChan
		switch cmd {
		case Read:
			dataChan <- state
		case Update:
			state = <-dataChan
		case Quit:
			slog.Info("KeepData - QUIT received, ending loop")
			break Loop
		default:
			slog.Warn(fmt.Sprintf("KeepData - Something else received: %v  -  %d", cmd, i))
		}
	}
	slog.Info("KeepData - Exit")
}

func GatherData(gatherDataCmdChan <-chan Command, keepDataCmdChan chan<- Command, dataChan chan *StateData, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	slog.Info("GatherData - Starting")
	for i := 0; ; i++ {
		if len(gatherDataCmdChan) > 0 {
			if <-gatherDataCmdChan == Quit {
				slog.Info("GatherData - QUIT received, ending loop")
				slog.Info("GatherData - Exit")
				return
			}
		}
		headTemp := i * 10
		state := Unknown
		if headTemp > 50 {
			state = Ready
		} else {
			state = Heating
		}

		keepDataCmdChan <- Update
		dataChan <- &StateData{headTemp, i * 3, state}
		time.Sleep(500 * time.Millisecond)
	}
}
