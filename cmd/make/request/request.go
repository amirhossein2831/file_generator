package request

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

// Request holds the template data
type Request struct {
	RequestName string
	FileName    string
}

// Convert PascalCase to snake_case
func toSnakeCase(str string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	return strings.ToLower(re.ReplaceAllString(str, "${1}_${2}"))
}

func CreateRequest(requestName string) {
	fileName := toSnakeCase(requestName)

	// Define template file and output directory
	tmplFile := "templates/request_template.go.tmpl"

	// get output dir
	outputDir := os.Getenv("PATH_FOR_REQUEST")
	if outputDir == "" {
		log.Fatalf("the path is not specific in env, please initiale the PATH_FOR_REQUEST variable")
	}

	// Parse the template
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Prepare template data
	req := Request{
		RequestName: requestName,
		FileName:    fileName,
	}

	// Ensure output directory exists

	if err = os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	// Output file path
	outputFile := filepath.Join(outputDir, fmt.Sprintf("%s.go", fileName))

	// Check if file already exists
	if _, err = os.Stat(outputFile); err == nil {
		log.Fatalf("Error: The file %s already exists.\n", outputFile)
	}

	// Create the output file
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	// Execute template and write to file
	err = tmpl.Execute(file, req)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	fmt.Printf("Request %s created successfully at %s\n", requestName, outputFile)
}

// RequestCmd generates a request file
var RequestCmd = &cobra.Command{
	Use:   "request [RequestName]",
	Short: "Create a new request struct file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		CreateRequest(args[0])
	},
}
