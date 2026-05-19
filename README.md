# ASCII Art Web

A web-based ASCII art generator built with Go. Type any text, choose a banner style, and instantly see your text rendered as ASCII art in the browser.

---

## Description

ASCII Art Web is an extension of the ASCII Art CLI project. It runs an HTTP server that serves a form where users can type text and select a banner style. The server processes the input and returns the rendered ASCII art directly in the browser.

The project supports all printable ASCII characters — from space (`32`) to tilde (`126`). Any character outside this range will return a `400 Bad Request` error.

**Available banner styles:**
- `standard`
- `shadow`
- `thinkertoy`

---

## Usage

**Requirements:** Go 1.18 or higher

```bash
git clone https://github.com/michaelbag8/ascii-art-web
cd ascii-art-web
go run .
```

Then open your browser and visit:

```
http://localhost:8080
```

1. Type your text into the input field
2. Select a banner style from the dropdown
3. Click **Generate ASCII Art**
4. Your ASCII art result appears below the form

---

## Implementation Details

The application is written in Go using only standard library packages. It works in three main stages:

### 1. Parsing Banner Files

Each banner file (`standard.txt`, `shadow.txt`, `thinkertoy.txt`) contains ASCII art representations of every printable character.

- Each character occupies **8 lines** of ASCII art followed by **1 empty separator line**
- Characters start at ASCII code `32` (space) and increment from there
- The parser reads the file and stores each character in a `map[rune][]string`, where the key is the character and the value is its 8 lines of art

### 2. Rendering Text Row by Row

ASCII art characters are rendered **horizontally** by printing matching rows across all characters simultaneously:

```
Row 1 of H + Row 1 of i
Row 2 of H + Row 2 of i
...
Row 8 of H + Row 8 of i
```

This produces properly aligned side-by-side ASCII art output.

### 3. Handling HTTP Requests

The server uses Go's `net/http` package with two endpoints:

| Method | Route        | Description                          |
|--------|-------------|--------------------------------------|
| GET    | `/`          | Serves the main HTML form            |
| POST   | `/ascii-art` | Processes input and returns ASCII art |

**Request flow:**

```
Browser → GET /
        ↓
Server serves HTML form
        ↓
User submits form → POST /ascii-art
        ↓
Server validates input
        ↓
Banner file loaded and parsed
        ↓
Text rendered row by row
        ↓
Result injected into HTML template
        ↓
Browser displays ASCII art
```

---

## Project Structure

```
ascii-art-web/
├── main.go
├── banner.go
├── render.go
├── generate.go
├── banner/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── templates/
│   ├── index.html
│   └── error.html
└── static/
    └── style.css
```

---

## Key Concepts

| Component      | Purpose                                    |
|----------------|--------------------------------------------|
| Banner file    | Stores ASCII character patterns            |
| ASCII map      | Associates characters with their art lines |
| Row rendering  | Combines matching rows horizontally        |
| GET `/`        | Serves the main webpage                    |
| POST `/ascii-art` | Processes and renders user input        |
| HTML template  | Displays generated ASCII art dynamically   |
| Error pages    | Returns meaningful 400, 404, 500 responses |

---

## Author

**Michael BAG**