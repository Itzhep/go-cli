# Go-CLI Tool

**My Go CLI Tool** is a command-line utility designed to simplify the setup of Go projects. With features for initializing new projects, setting up Git repositories, and choosing from various project templates, it streamlines your development workflow.

## Features

- Create new Go projects with multiple templates (basic, web-server, CLI tool).
- Optionally initialize a Git repository.
- Configure projects using a JSON file or interactive prompts.

## Installation

You can install **My Go CLI Tool** on different platforms as follows:

### Pre-built Binaries

1. **Download the Binary**:
   - Visit the [Releases page](https://github.com/Itzhep/go-cli/releases) to download the appropriate binary for your operating system.

2. **Install the Binary**:

   - **For Linux/macOS**:
     ```bash
     curl -LO https://github.com/Itzhep/go-cli/releases/download/v1.0.0/go-cli.exe
     chmod +x my-go-cli
     sudo mv my-go-cli /usr/local/bin/
     ```

   - **For Windows**:
     ```powershell
     curl -LO https://github.com/Itzhep/my-go-cli/releases/download/v1.0.0/go-cli.exe
     move go-cli.exe C:\path\to\your\bin
     ```
     📓 Note : if it dosnt work run it from src by :
     ```bash
      go run Go-CLI.go
     ```
 3. **Configuration via JSON**
You can also configure your project setup using a JSON configuration file. Create a file with the following format:

```json

{
  "projectName": "my-project",
  "gitInit": true,
  "template": "cli-tool"
}
```
Then, run:
```bash
my-go-cli --config path/to/config.json
```
# License
My Go CLI Tool is released under the MIT License.
