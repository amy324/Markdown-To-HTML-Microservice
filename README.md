# Markdown to HTML Converter

A lightweight **Markdown-to-HTML converter written in Go**.

The project provides a `MarkdownToHTML` function that takes Markdown content and converts common Markdown syntax into HTML using regular expressions and the Go standard library.

## Features

* Convert headings from `#` to `######`
* Convert **bold** and *italic* text
* Convert unordered lists using `*` or `-`
* Convert ordered lists
* Convert Markdown links
* Convert Markdown images
* No external dependencies

## Installation

Make sure Go is installed, then clone the repository:

```bash
git clone https://github.com/amy324/Markdown-To-HTML-Microservice.git
cd Markdown-To-HTML-Microservice
```

## Usage

Pass Markdown content to the `MarkdownToHTML` function:

```go
markdown := `
# Markdown to HTML Converter

This is a **simple** example of _Markdown_ to HTML conversion.

## Lists

* Item 1
* Item 2

1. Numbered item 1
2. Numbered item 2

## Links

[Visit Google](https://www.google.com)

## Images

![Markdown Logo](https://markdown-here.com/img/icon256.png)
`

html := MarkdownToHTML(markdown)
fmt.Println(html)
```

Example output:

```html
<h1>Markdown to HTML Converter</h1>

This is a <strong>simple</strong> example of <em>Markdown</em> to HTML conversion.

<h2>Lists</h2>

<ul>
<li>Item 1</li>
<li>Item 2</li>
</ul>

<ol>
<li>Numbered item 1</li>
<li>Numbered item 2</li>
</ol>

<h2>Links</h2>

<a href="https://www.google.com">Visit Google</a>

<h2>Images</h2>

<img src="https://markdown-here.com/img/icon256.png" alt="Markdown Logo">
```

## How It Works

The conversion is handled by the `MarkdownToHTML` function in `main.go`.

Regular expressions are used to identify different Markdown patterns and replace them with the corresponding HTML.

For example:

* `# Heading` → `<h1>Heading</h1>`
* `**bold**` → `<strong>bold</strong>`
* `_italic_` → `<em>italic</em>`
* `* Item` → `<ul><li>Item</li></ul>`
* `1. Item` → `<ol><li>Item</li></ol>`
* `[text](url)` → `<a href="url">text</a>`
* `![text](url)` → `<img src="url" alt="text">`

The project uses Go's `regexp` and `strings` packages for the parsing and transformation.

## Project Structure

```text
main.go
```

The `MarkdownToHTML` function contains the conversion logic, while `main()` provides an example of how to use it.

## Dependencies

This project uses only the **Go standard library** and does not require any external dependencies.

## Purpose

This project is a useful exercise in working with regular expressions, string processing, and translating a set of syntax rules into a reusable Go function.

## Contributing

Contributions, suggestions, and bug reports are welcome. Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
