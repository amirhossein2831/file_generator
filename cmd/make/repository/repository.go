package repository

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

//go:embed repository_template.go.tmpl
var templates embed.FS

// Repository holds the template data
type Repository struct {
	RepositoryName string
	FileName       string
}

// Convert PascalCase to snake_case
func toSnakeCase(str string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	return strings.ToLower(re.ReplaceAllString(str, "${1}_${2}"))
}

func CreateRepository(repositoryName string) {
	// Convert repository name to snake_case for the file name
	fileName := toSnakeCase(repositoryName)

	// Define the template file path (relative to the package's embedded filesystem)
	tmplFile := "repository_template.go.tmpl" // File name for the embedded template

	// Get output directory from the environment variable
	outputDir := os.Getenv("PATH_FOR_REPOSITORY")
	if outputDir == "" {
		log.Fatalf("the path is not specific in env, please initialize the PATH_FOR_REPOSITORY variable")
	}

	// Parse the template from the embedded filesystem
	tmpl, err := template.ParseFS(templates, tmplFile)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Prepare the data for the template
	repo := Repository{
		RepositoryName: repositoryName,
		FileName:       fileName,
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

	// Execute template and write to file
	err = tmpl.Execute(file, repo)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	fmt.Printf("Repository %s created successfully at %s\n", repositoryName, outputFile)

}

// RepositoryCmd generates a repository file
var RepositoryCmd = &cobra.Command{
	Use:   "repository [RepositoryName]",
	Short: "Create a new repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		CreateRepository(args[0])
	},
}
