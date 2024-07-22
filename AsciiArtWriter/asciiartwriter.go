package asciiartwriter

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	printasciiart "github.com/ahatdemirezen/docker/PrintAsciiArt"
)

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
	res := printasciiart.PrintAsciiArt(sepArgs, lines)
	return res
}
