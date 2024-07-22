package asciitoweb

import (
	"fmt"
	"net/http"
	"os"

	asciiartwriter "github.com/ahatdemirezen/docker/AsciiArtWriter"
)

func AsciiToWeb(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "templates/form.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		str := r.FormValue("str")
		banner := r.FormValue("banner")
		if str == "" {
			str = "String Is Empty!"
		}
		_, err := os.Open("fonts/" + banner + ".txt")
		if err != nil {
			fmt.Fprintf(w, "%s\n", "Banner is not exist!")
			return
		}
		for _, char := range str {
			if char >= 1 && char <= 127 {
				continue
			} else {
				str = "Invalid Character."
			}
		}
		res := asciiartwriter.AsciiArtWriter(str, banner)
		fmt.Fprintf(w, "%s\n", res)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
