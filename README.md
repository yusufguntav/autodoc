# Auto-Document

Auto-Document is a Go package that automatically generates documentation for your project based on git commits using AI. It analyzes your unpushed Git commits and generates comprehensive documentation automatically.

## Features

- Analyzes unpushed Git commits in your current branch
- Uses Gemini AI to generate human-readable documentation
- Easy to integrate into your development workflow

## Installation

### As a CLI tool

```bash
go install github.com/yusufguntav/Auto-Document/cmd/auto-doc@latest
```

### As a library

```bash
go get github.com/yusufguntav/Auto-Document
```

## Prerequisites

- Go 1.18 or later
- Git installed and accessible in your PATH
- A Gemini API key

## Configuration

Create a `.env` file in your project root with the following content:

```
GEMINI_API_KEY=your_gemini_api_key_here
```

## Usage

### CLI Usage

After installation, run:

```bash
auto-doc
```

This will analyze your unpushed commits and generate documentation.

### Library Usage

```go
package main

import (
	"fmt"
	
	"github.com/joho/godotenv"
	"github.com/yusufguntav/Auto-Document"
)

func main() {
	// Load environment variables from .env file
	godotenv.Load()
	
	// Get current branch
	branch, err := autodocument.GetCurrentBranch()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	// Get commits from branch
	commits, err := autodocument.GetCommits(branch)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	// Generate documentation with AI
	documentation := autodocument.SendMessageAI(commits)
	fmt.Println(documentation)
}
```

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
