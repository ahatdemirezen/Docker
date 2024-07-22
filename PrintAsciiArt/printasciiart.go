package printasciiart

func PrintAsciiArt(sentences []string, textFile []string) string {
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
