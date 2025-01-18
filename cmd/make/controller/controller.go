package controller

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

// Controller to hold template data
type Controller struct {
	ControllerName string
	FileName       string
}

// Function to convert PascalCase to snake_case
func toSnakeCase(str string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(snake)
}

func CreateController(controllerName string) {
	// Convert to snake_case for the file name
	fileName := toSnakeCase(controllerName)

	// Define the template file
	tmplFile := "templates/controller_template.go.tmpl"

	// Parse the template
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Prepare the data for the template
	controller := Controller{
		ControllerName: controllerName,
		FileName:       fileName,
	}

	// Define the output file path
	outputDir := os.Getenv("PATH_FOR_CONTROLLER")
	if outputDir == "" {
		log.Fatalf("the path is not specific in env, please initiale the PATH_FOR_CONTROLLER variable")
	}

	if err = os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Error creating directories: %v", err)
	}

	// Output file path
	outputFile := filepath.Join(outputDir, fmt.Sprintf("%s.go", fileName))

	// Check if the file already exists
	if _, err = os.Stat(outputFile); err == nil {
		log.Fatalf("Error: The file %s already exists.\n", outputFile)
	}

	// Create the output file
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	// Execute the template and write to the file
	err = tmpl.Execute(file, controller)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	// Success message
	fmt.Printf("Controller %s created successfully at %s\n", controllerName, outputFile)
}

// ControllerCmd for generating the controller file
var ControllerCmd = &cobra.Command{
	Use:   "controller [ControllerName]",
	Short: "Create a new controller",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		CreateController(args[0])
	},
}
