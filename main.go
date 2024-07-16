package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func printAsciiArt(sentences []string, textFile []string) string {
	res := ""
	for i, word := range sentences {
		if word == "" {
			if i != 0 {
				res += "\n"
			}
			continue
		}
		lineLen := 0
		for i := 0; i < len(word); i++ {
			for lineIndex, line := range textFile {
				if lineIndex == (int(word[i])-32)*9+2 {
					lineLen += len(line)
				}
			}
		}
		if lineLen > 177 {
			word = "Input Is Too Long!"
		}
		for h := 1; h < 9; h++ {
			for i := 0; i < len(word); i++ {
				for lineIndex, line := range textFile {
					if lineIndex == (int(word[i])-32)*9+h {
						res += line
					}
				}
			}
			res += "\n"
		}
	}
	return res
}

func AsciiArtWriter(argStr string, bannerType string) string {
	sepArgs := strings.Split(argStr, "\n")
	file, err := os.Open("fonts/" + bannerType + ".txt")
	if err != nil {
		fmt.Println("Error")
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}
	res := printAsciiArt(sepArgs, lines)
	return res
}

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
		res := AsciiArtWriter(str, banner)
		fmt.Fprintf(w, "%s\n", res)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.Handle("/templates/style.css", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	http.HandleFunc("/", AsciiToWeb)
	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Something went wrong!")
		return
	}
}
