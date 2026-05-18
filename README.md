# 🧮 Ruach_Calculator

Ruach_Calculator is a simple GUI-based calculator built with Go using the Fyne framework.  
It supports basic arithmetic operations as well as advanced features like square root, exponentiation, and percentage calculations.

---

## 🚀 Features

- Basic operations: `+`, `-`, `x`, `÷`
- Square root: `√`
- Power operation: `^`
- Percentage calculation: `%`
- Decimal support
- Delete (DEL) and Clear (C) functions
- Real-time expression display
- Error handling for invalid expressions
- Built with a graphical interface using Fyne

---

## 🛠️ Tech Stack

- Go (Golang)
- Fyne GUI Toolkit

---

## 📦 Installation

### 1. Clone the repository
```bash
git clone https://github.com/your-username/Ruach_Calculator.git
cd Ruach_Calculator
```
### 2. Install dependencies
```bash
go mod tidy
```
### 3. Run the application
```bash
go run main.go
```

---

#### 🧠 How It Works

The calculator:

- Takes input as a string slice
- Parses numbers and operators
- Processes operations in order of precedence:
- square root (√)
- Power (^)
- Percentage (%)
- Multiplication and Division (× and ÷)
- Addition and Subtraction (+ and −)

Returns a final computed result

## 🖥️ Build Executable

Linux / macOS
```bash
go build -o Ruach_Calculator
```

Windows
```bash
GOOS=windows GOARCH=amd64 go build -o Ruach_Calculator.exe
```
---

## 📁 Project Structure
```
Ruach_Calculator
├── main.go
├── fyne-cross
    └── bin/windows-amd64
        └── Ruach Calculator.exe
    └── dist/windows-amd64
        └── Ruach Calculator.zip
    └── temp/windows-amd64
        └── Icon.png
├── go.mod
├── go.sum
└── README.md
```
---

## ⚠️ Notes

- This project uses manual expression parsing (no external math parser library).
- Invalid expressions return an error message on screen.
- GUI is built using the Fyne framework.

---

## 💡 Future Improvements

- Add parentheses support ()
- Add history of calculations
- Improve error handling system
- Add scientific functions (sin, cos, tan, log)

---

## 👨🏽‍💻 Author
```
Built by Ayomide Shogbesan (Jedidiah)
Passion: Data, Cloud Engineering, and Software Systems 🚀
```
---

## 📜 License

This project is open-source and free to use.