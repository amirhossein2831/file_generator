package make

import (
	"file_generator/cmd/make/all"
	"file_generator/cmd/make/controller"
	"file_generator/cmd/make/exception"
	"file_generator/cmd/make/repository"
	"file_generator/cmd/make/request"
	"file_generator/cmd/make/service"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"time"
)

// Make Commands for making file
var Make = &cobra.Command{
	Use:   "make",
	Short: "Commands for making file",
}

func init() {
	// init requirement
	beforeMake()

	// init subcommand
	Make.AddCommand(
		exception.ExceptionCmd,
		controller.ControllerCmd,
		service.ServiceCmd,
		repository.RepositoryCmd,
		request.RequestCmd,
		all.AllCmd,
	)
}

func beforeMake() {
	// Initialize Env variable
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ENV Service: Failed to  loading .env file. %v.    timestamp: %s", err, time.Now().String())
	}
}
