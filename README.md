# Auto-Document

**Auto-Document** is a Go package that automatically generates documentation for your project by analyzing your unpushed Git commits using AI (Gemini). This helps keep your documentation up-to-date with minimal effort.

---

## 🚀 Features

- 🔍 Analyzes **unpushed Git commits** in your current branch  
- 🤖 Uses **Gemini AI** to generate clear, human-readable documentation  
- 🔧 Easily integrates into your existing development workflow

---

## 📦 Installation

Install the package using:

```bash
go get github.com/yusufguntav/autodoc
```
✅ Prerequisites
Before using Auto-Document, ensure the following:

Go 1.18+ installed

Git installed and accessible via command line

A valid Gemini API key

⚙️ Configuration
Create a .env file in the root of your project with:

GEMINI_API_KEY=your_gemini_api_key_here
🧠 Usage
```
package main

import (
	"github.com/yusufguntav/autodoc"
)

func main() {
	// Your application logic

	autodoc.GenerateDocumentation()

	// Your application logic
}
```
Start your project using the autodoc command to trigger documentation generation.

```
// Generate document (print terminal)
go run . autodoc

// Generate document to specific file
go run . autodoc -o document.txt
```

📄 License
This project is licensed under the MIT License.

🤝 Contributing
Contributions are welcome! Feel free to submit a pull request or open an issue.
