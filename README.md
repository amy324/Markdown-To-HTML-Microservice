
# Markdown to HTML Microservice

This is a microservice written in Golang that converts Markdown syntax to HTML. It provides a function `MarkdownToHTML` that takes Markdown content as input and returns the corresponding HTML content.

## Usage

To use this microservice, you can call the `MarkdownToHTML` function with the Markdown content as an argument. It will then convert the Markdown to HTML and return the HTML content.

Example usage:

```go
func main() {
	markdown := `
# Markdown to HTML Converter

This is a **simple** example of _Markdown_ to HTML conversion.

## Headings

### This is a subheading

#### This is a sub-subheading

## Lists

* Item 1
* Item 2

1. Numbered item 1
2. Numbered item 2

## Links

[Visit Google](https://www.google.com)

## Images

[Markdown Logo](https://markdown-here.com/img/icon256.png)
`
	html := MarkdownToHTML(markdown)
	fmt.Println(html)
}
```

Output:

```html

<h1>Markdown to HTML Converter</h1>


This is a <strong>simple</strong> example of <em>Markdown</em> to HTML conversion.

<h2>Headings</h2>


<h3>This is a subheading</h3>


<h4>This is a sub-subheading</h4>


<h2>Lists</h2>


<ul>
<li> Item 1</li>
<li> Item 2</li>
</ul>

<ol>
<li> Numbered item 1</li>
<li> Numbered item 2</li>
</ol>

<h2>Links</h2>


<a href="https://www.google.com">Visit Google</a>

<h2>Images</h2>


<a href="https://markdown-here.com/img/icon256.png">Markdown Logo</a>

```

## Installation

1. Install Go if you haven't already: [Go Installation Guide](https://golang.org/doc/install)

2. Add the package to your project using Go modules:

```bash
go get github.com/amy324/markdown-to-html-microservice
```

3. Import the package in your Go code and use the `MarkdownToHTML` function as shown in the example above.

## Code Explanation

### `markdown.go`

The `markdown.go` file contains the logic for converting Markdown to HTML. It defines the `MarkdownToHTML` function, which takes Markdown content as input and uses regular expressions to identify and replace Markdown syntax with corresponding HTML tags. This includes handling headings, bold and italic text, unordered and ordered lists, links, and images.


### Regular Expressions

Regular expressions (regex) are patterns used to match character combinations in strings. They are widely used for text search and manipulation tasks. In this microservice, regex is utilised extensively to identify and transform Markdown elements into HTML tags.

### Logic Implementation

#### Headings

Regex is used to identify lines starting with `#` and replace them with corresponding HTML heading tags (`<h1>` to `<h6>`).

#### Bold and Italics

Similarly, regex is used to identify patterns for bold (`**text**`) and italic (`_text_`) text and replace them with `<strong>` and `<em>` HTML tags, respectively.

#### Lists

For unordered and ordered lists, regex is used to match lines starting with `*` or a number followed by `.`. These lines are then wrapped in `<ul>` and `<ol>` tags, respectively.

#### Links and Images

Regex patterns are used to identify Markdown link and image syntax (`[text](url)` and `![text](url)`) and replace them with corresponding HTML `<a>` and `<img>` tags.

### Conclusion

Regular expressions provide a powerful way to parse and manipulate text efficiently. By leveraging regex patterns, we can easily convert Markdown syntax into HTML markup, making this microservice an efficient tool for text conversion tasks.


## Dependencies

This microservice does not have any external dependencies beyond the Golang standard library.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License 
