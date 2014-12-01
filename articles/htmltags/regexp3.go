package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := `<body>Text "..." '...'
                <input type="button" onclick="alert(2>1)">
                </body>
                <a href="www<>.exa<mple.com"target="_>blank"></a>
                <script>
                function text(node) {
                        for (var i = 0; i < node.childNodes.length; i++) {
                                var n = node.childNodes[i];
                        }
                }
                </script>`

	r := regexp.MustCompile(
                `(?s:(<style>|<script>).+</(style|script)>|` + 
                `<(\s*[^/][^\s>]*)(=\s*("[^"]*"|'[^']*'|[^\s>]+)|[^>])*(>|$))`)
	
        for _, submatch := range r.FindAllStringSubmatch(text, -1) {
		switch {
		// <script> and <style>.
		case submatch[1] != "":
			fmt.Println(submatch[1])
		// Everything else.
		default:
			fmt.Println(submatch[0])
		}
	}
}
