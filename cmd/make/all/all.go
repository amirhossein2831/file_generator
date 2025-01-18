package all

import (
	"github.com/amirhossein2831/file_generator/cmd/make/controller"
	"github.com/amirhossein2831/file_generator/cmd/make/exception"
	"github.com/amirhossein2831/file_generator/cmd/make/repository"
	"github.com/amirhossein2831/file_generator/cmd/make/request"
	"github.com/amirhossein2831/file_generator/cmd/make/service"
	"github.com/spf13/cobra"
)

// AllCmd for generating the all file
var AllCmd = &cobra.Command{
	Use:   "all [AllName]",
	Short: "Create a new file for all module",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		controller.CreateController(args[0] + "Controller")
		request.CreateRequest(args[0] + "Request")
		repository.CreateRepository(args[0] + "Repository")
		service.CreateService(args[0] + "Service")
		exception.CreateException(args[0] + "Exception")
	},
}
