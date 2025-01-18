package service

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

// Service to hold template data
type Service struct {
	ServiceName string
	FileName    string
}

// Function to convert PascalCase to snake_case
func toSnakeCase(str string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(snake)
}

func CreateService(serviceName string) {
	// Convert to snake_case for the file name
	fileName := toSnakeCase(serviceName)

	// Define the template file
	tmplFile := "templates/service_template.go.tmpl"

	// Parse the template
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Prepare the data for the template
	service := Service{
		ServiceName: serviceName,
		FileName:    fileName,
	}

	// Define the output file path
	outputDir := os.Getenv("PATH_FOR_SERVICE")
	if outputDir == "" {
		log.Fatalf("the path is not specific in env, please initiale the PATH_FOR_SERVICE variable")
	}

	// make file
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
	err = tmpl.Execute(file, service)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	// Success message
	fmt.Printf("Service %s created successfully at %s\n", serviceName, outputFile)
}

// ServiceCmd for generating the service file
var ServiceCmd = &cobra.Command{
	Use:   "service [ServiceName]",
	Short: "Create a new service",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		CreateService(args[0])
	},
}
