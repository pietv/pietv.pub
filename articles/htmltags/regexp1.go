package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	text := `<!doctype html>
		<!-- comment -->
		<html>>
		<head data-one="valu'e" data-two='valu"e'></head>
		<body>Text
		</body>
		<p><b></p></b>
		< br >
		<hr/>
		<a href="www.example.com"target="_blank"></a>
		</html>>`

	r := regexp.MustCompile(`(?Us:<[^/>].*>)`)
	fmt.Println(strings.Join(r.FindAllString(text, -1), "\n"))
}
