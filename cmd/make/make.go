package make

import (
	"github.com/amirhossein2831/file_generator/cmd/make/all"
	"github.com/amirhossein2831/file_generator/cmd/make/controller"
	"github.com/amirhossein2831/file_generator/cmd/make/exception"
	"github.com/amirhossein2831/file_generator/cmd/make/repository"
	"github.com/amirhossein2831/file_generator/cmd/make/request"
	"github.com/amirhossein2831/file_generator/cmd/make/route"
	"github.com/amirhossein2831/file_generator/cmd/make/service"
	"github.com/spf13/cobra"
)

// Make Commands for making file
var Make = &cobra.Command{
	Use:   "make",
	Short: "Commands for making file",
}

func init() {
	// init requirement
	//beforeMake()

	// init subcommand
	Make.AddCommand(
		exception.ExceptionCmd,
		controller.ControllerCmd,
		service.ServiceCmd,
		repository.RepositoryCmd,
		request.RequestCmd,
		route.RouteCmd,
		all.AllCmd,
	)
}

//func beforeMake() {
//	// Initialize Env variable
//	err := godotenv.Load()
//	if err != nil {
//		err = godotenv.Load("env/development.env")
//
//		log.Fatalf("ENV Service: Failed to  loading .env file. %v.    timestamp: %s", err, time.Now().String())
//	}
//
//}
