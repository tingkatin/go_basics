package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func htmlParse(n *html.Node) (words, images int) {

	if n.Type == html.TextNode {
		parent := n.Parent
		if parent != nil && (parent.Data == "script" || parent.Data == "style") {
			return
		}
		words += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		w, i := htmlParse(child)

		words += w
		images += i
	}
	return
}

func main() {

	response, err := http.Get("http://google.com")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot fetch data: %s\n", err)
		os.Exit(-1)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading body: %s\n", err)
		os.Exit(-1)
	}

	doc, err := html.Parse(bytes.NewReader([]byte(body)))

	if err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %s\n", err)
		os.Exit(-1)
	}

	words, pics := htmlParse(doc)

	fmt.Printf("%d words and %d images\n", words, pics)
}
