package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Closable interface {
	// Close release all resources used by this object, including goroutines.
	Close() error
}

type myProcessor struct {
	stopProcessing bool
}

func (m *myProcessor) Process() {
	go func() {
		for !m.stopProcessing {
			fmt.Println("Processing...")
			time.Sleep(time.Second)
		}
		fmt.Println("Processor has stopped")
	}()
}

func (m *myProcessor) Close() error {
	// Close should only return when it's certain that the for loop
	// in the Process function has stopped.
	return nil
}

func main() {
	myProcessor := &myProcessor{}
	myProcessor.Process()
	closables := []Closable{myProcessor}

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// In case we are running in a containerised environment there are different reasons
	// why our program can receive an os signal to terminate. This might for example
	// during a rolling update of the container.
	// If our program performs some processing, for example consuming a stream, we will need
	// to make sure our program exits gracefully. I.e. ensure it finishes up it's last task
	// and stops with whatever it was doing.
	// Implement the close function such that it stops the for loop and waits until it has
	// actually stopped. Hint: use a channel to do this.
	sig := <-sigs
	fmt.Println("Received:", sig)
	for _, closable := range closables {
		err := closable.Close()
		if err != nil {
			fmt.Println("unable to gracefully exit: ", err)
		}
	}
	fmt.Println("graceful exit!")
}
