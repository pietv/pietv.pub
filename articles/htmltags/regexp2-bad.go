package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	text := `<input type="button" onclick="alert(2>1)">
		<a href="www<>.exa<mple.com"target="_>blank"></a>
		<script>
		function text(node) {
			for (var i = 0; i < node.childNodes.length; i++) {
		   		var n = node.childNodes[i];
			}
		}
		</script>`

	r := regexp.MustCompile(
		`(?s:<(\s*[^/][^\s>]*)(=\s*("[^"]*"|'[^']*'|[^\s>]+)|[^>])*(>|$))`)
	fmt.Println(strings.Join(r.FindAllString(text, -1), "\n"))
}
