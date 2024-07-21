package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	version = "1.0.0"
	rootCmd  = &cobra.Command{
		Use:   "go-cli",
		Short: "A CLI tool for setting up Go projects",
		Run:   run,
	}
	configFile string
)


type ProjectConfig struct {
	ProjectName string `json:"projectName"`
	GitInit     bool   `json:"gitInit"`
	Template    string `json:"template"`
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Show version")
	rootCmd.Flags().BoolP("help", "h", false, "Show help")
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "", "Path to a JSON configuration file")
}

func run(cmd *cobra.Command, args []string) {
	showVersion, _ := cmd.Flags().GetBool("version")
	if showVersion {
		fmt.Printf("go-cli version %s\n", version)
		return
	}

	showHelp, _ := cmd.Flags().GetBool("help")
	if showHelp {
		cmd.Help()
		return
	}

	// Check for updates
	err := checkForUpdates("go-cli", version)
	if err != nil {
		fmt.Println("Error checking for updates:", err)
		return
	}

	var answers ProjectConfig

	if configFile != "" {
		config, err := readConfig(configFile)
		if err != nil {
			fmt.Println("Error reading configuration file:", err)
			return
		}
		answers = config
	} else {
		prompts := []*survey.Question{
			{
				Name:     "projectName",
				Prompt:   &survey.Input{Message: "What is the project name?"},
				Validate: survey.Required,
			},
			{
				Name:   "gitInit",
				Prompt: &survey.Confirm{Message: "Do you want to initialize a Git repository?", Default: true},
			},
			{
				Name: "template",
				Prompt: &survey.Select{
					Message: "Choose a project template:",
					Options: []string{"basic", "web-server", "cli-tool"},
					Default: "basic",
				},
			},
		}
		err = survey.Ask(prompts, &answers)
		if err != nil {
			fmt.Println("Error asking questions:", err)
			return
		}
	}

	err = createGoProjectStructure(answers.ProjectName, answers.GitInit, answers.Template)
	if err != nil {
		fmt.Println("Error during project setup:", err)
	}
}

func readConfig(filePath string) (ProjectConfig, error) {
	var config ProjectConfig
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	return config, err
}

func checkForUpdates(packageName, currentVersion string) error {
	fmt.Println(color.YellowString("Checking for updates is not implemented in this Go version."))
	return nil
}

func execCommand(command, cwd string) error {
	cmd := exec.Command("cmd", "/C", command)
	cmd.Dir = cwd
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing command: %s, output: %s", err, string(output))
	}
	return nil
}

func createGoProjectStructure(projectName string, gitInit bool, template string) error {
	projectPath := filepath.Join(".", projectName)

	if _, err := os.Stat(projectPath); !os.IsNotExist(err) {
		return fmt.Errorf("directory %s already exists", projectPath)
	}

	err := os.MkdirAll(projectPath, os.ModePerm)
	if err != nil {
		return err
	}

	spinner := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	spinner.Start()
	defer spinner.Stop()

	err = execCommand("go mod init "+projectName, projectPath)
	if err != nil {
		return err
	}

	switch template {
	case "web-server":
		createWebServerTemplate(projectPath)
	case "cli-tool":
		createCLIToolTemplate(projectPath)
	default:
		createBasicTemplate(projectPath)
	}

	if gitInit {
		err = execCommand("git init", projectPath)
		if err != nil {
			color.Yellow("Git is not installed or initialization failed.")
		} else {
			color.Green("Git repository initialized.")
		}
	}

	readmeContent := fmt.Sprintf("# %s\n\nYour project description here.", projectName)
	err = ioutil.WriteFile(filepath.Join(projectPath, "README.md"), []byte(readmeContent), os.ModePerm)
	if err != nil {
		return err
	}

	color.Green(fmt.Sprintf("Go project %s has been created successfully!", projectName))
	return nil
}

func createBasicTemplate(projectPath string) error {
	mainGoContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
`
	return ioutil.WriteFile(filepath.Join(projectPath, "main.go"), []byte(mainGoContent), os.ModePerm)
}

func createWebServerTemplate(projectPath string) error {
	mainGoContent := `package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, web server!"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
`
	return ioutil.WriteFile(filepath.Join(projectPath, "main.go"), []byte(mainGoContent), os.ModePerm)
}

func createCLIToolTemplate(projectPath string) error {
	mainGoContent := `package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "world", "name to greet")
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
`
	return ioutil.WriteFile(filepath.Join(projectPath, "main.go"), []byte(mainGoContent), os.ModePerm)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
