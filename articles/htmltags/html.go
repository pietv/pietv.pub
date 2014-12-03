package main

import (
	"fmt"
	"strings"

	"code.google.com/p/go.net/html"
)

func main() {
	text := `<!doctype html><body>Text "..." '...'
		<!-- -->
		<input type="button" onclick="alert(2>1)">
		</body>
		<a href="www<>.exa<mple.com"target="_>blank"></a>
		< br ><hr/>
		<script>
		function text(node) {
			for (var i = 0; i < node.childNodes.length; i++) {
				var n = node.childNodes[i];
			}
		}
		</script>`

	doc := html.NewTokenizer(strings.NewReader(text))
	for {
		token := doc.Next()
		if token == html.ErrorToken {
			break
		}
		switch token {
		case html.StartTagToken, html.SelfClosingTagToken:
			fmt.Println(doc.Token())
		}
	}
}
