package command

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/jngreenwood/nazgul/internal/probe"
	"github.com/jngreenwood/nazgul/internal/server"
	"github.com/jngreenwood/nazgul/internal/worker"
	"github.com/mitchellh/cli"
)

type RunCommand struct {
	Ui cli.Ui
}

func (f *RunCommand) Help() string {
	helpText := `
	Usage: nazgul run [options] [args]
	`
	return strings.TrimSpace(helpText)
}

func (f *RunCommand) Synopsis() string {
	return "Todo"
}

func (f *RunCommand) Name() string { return "run" }

func (f *RunCommand) Run(args []string) int {

	ctx, cancel := context.WithCancel(context.Background())

	//create the webserver
	go server.Serve()

	scheduler := worker.NewJobScheduler(ctx)
	go scheduler.Start()

	job := &probe.TestJob{
		Message: "Hello World",
	}

	pingjob := &probe.PingJob{
		Host: "192.168.1.1",
	}
	scheduler.AddTask(job)
	scheduler.AddTask(pingjob)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received.
	fmt.Println("Press Ctrl+C to stop")
	s := <-sigChan
	fmt.Printf("Signal received: %v, stopping...\n", s)
	cancel()

	fmt.Println("Program exited gracefully")

	return 0

}
