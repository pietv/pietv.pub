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
	           <head data-one="valu'e>" data-two='valu"e>'></head>
	           <<body>Text "..." '...'
	           <input type="button" onclick="alert(2>1)">
                   </body>
                   <p><b></p></b>
                   < br >
                   <hr/>
                   <a href="www<>.exa<mple.com"target="_>blank"></a>
	         </html>>`

	r := regexp.MustCompile(`(?Us:<[^/>].*>)`)
	fmt.Println(strings.Join(r.FindAllString(text, -1), "\n"))
}
