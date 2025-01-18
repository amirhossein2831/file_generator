package service

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed service_template.go.tmpl
var templates embed.FS

// Service holds the template data
type Service struct {
	ServiceName string
	FileName    string
}

// Convert PascalCase to snake_case
func toSnakeCase(str string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	return strings.ToLower(re.ReplaceAllString(str, "${1}_${2}"))
}

func CreateService(serviceName string) {
	// Convert service name to snake_case for the file name
	fileName := toSnakeCase(serviceName)

	// Define the template file path (relative to the embedded filesystem)
	tmplFile := "service_template.go.tmpl" // File name for the embedded template

	// Get output directory from the environment variable
	outputDir := os.Getenv("PATH_FOR_SERVICE")
	if outputDir == "" {
		log.Fatalf("The path is not specified in env, please initialize the PATH_FOR_SERVICE variable")
	}

	// Parse the template from the embedded filesystem
	tmpl, err := template.ParseFS(templates, tmplFile)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Prepare the data for the template
	service := Service{
		ServiceName: serviceName,
		FileName:    fileName,
	}

	// Ensure the output directory exists
	if err = os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Error creating directories: %v", err)
	}

	// Define the output file path
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
	err = tmpl.Execute(file, service)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	// Success message
	fmt.Printf("Service %s created successfully at %s\n", serviceName, outputFile)
}

// ServiceCmd generates the service file
var ServiceCmd = &cobra.Command{
	Use:   "service [ServiceName]",
	Short: "Create a new service",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		CreateService(args[0])
	},
}
