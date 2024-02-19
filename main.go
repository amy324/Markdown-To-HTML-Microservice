package main

import (
	"fmt"
	"regexp"
	"strings"
)

// MarkdownToHTML converts Markdown to HTML
func MarkdownToHTML(markdown string) string {
	// Headings
	markdown = regexp.MustCompile(`(?m)^###### (.*)$`).ReplaceAllString(markdown, "<h6>$1</h6>\n")
	markdown = regexp.MustCompile(`(?m)^##### (.*)$`).ReplaceAllString(markdown, "<h5>$1</h5>\n")
	markdown = regexp.MustCompile(`(?m)^#### (.*)$`).ReplaceAllString(markdown, "<h4>$1</h4>\n")
	markdown = regexp.MustCompile(`(?m)^### (.*)$`).ReplaceAllString(markdown, "<h3>$1</h3>\n")
	markdown = regexp.MustCompile(`(?m)^## (.*)$`).ReplaceAllString(markdown, "<h2>$1</h2>\n")
	markdown = regexp.MustCompile(`(?m)^# (.*)$`).ReplaceAllString(markdown, "<h1>$1</h1>\n")

	// Bold
	markdown = regexp.MustCompile(`\*\*(.*?)\*\*`).ReplaceAllString(markdown, "<strong>$1</strong>")

	// Italics
	markdown = regexp.MustCompile(`(?m)(^|[^_])_([^_]+)_($|[^_])`).ReplaceAllString(markdown, "$1<em>$2</em>$3")

	// Unordered Lists
	markdown = regexp.MustCompile(`(?m)((?:^[\*\-].*$\s*){2,})`).ReplaceAllStringFunc(markdown, func(s string) string {
		lines := strings.Split(strings.TrimSpace(s), "\n")
		listItems := ""
		for _, line := range lines {
			listItems += "<li>" + line[1:] + "</li>\n"
		}
		return "<ul>\n" + listItems + "</ul>\n"
	})

	// Ordered Lists
	markdown = regexp.MustCompile(`(?m)((?:^\d+\..*$\s*){2,})`).ReplaceAllStringFunc(markdown, func(s string) string {
		lines := strings.Split(strings.TrimSpace(s), "\n")
		listItems := ""
		for _, line := range lines {
			listItems += "<li>" + line[2:] + "</li>\n"
		}
		return "<ol>\n" + listItems + "</ol>\n"
	})

	// Links
	markdown = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`).ReplaceAllString(markdown, "<a href=\"$2\">$1</a>")

	// Images
	markdown = regexp.MustCompile(`!\[([^\]]+)\]\(([^)]+)\)`).ReplaceAllString(markdown, "<img src=\"$2\" alt=\"$1\">")

	return markdown
}

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
