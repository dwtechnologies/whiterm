package whiterm

// Remove removes all spaces and tabs from String s.
// It will also remove everything regarded as a comment from s (#).
// It supports  "\"-escaping and doesn't remove anything inside of quotes.
// It returns the formated String.
func Remove(s string) string {
	newtext := ""
	escapeCount := 0

	// Loop through the s
	for pos := 0; pos < len(s); pos++ {
		char := string(s[pos])

		// Find escape character and skip parsing the char after
		if char == "\\" {
			char := string(s[pos+1])
			newtext += char
			pos++
			continue
		}

		// Find Quote and indicate that we shouldn't remove spaces or tabs until escapeCount is even
		if char == "\"" {
			escapeCount++
			continue
		}

		// Find comments
		if char == "#" {
			if escapeCount%2 == 1 {
				newtext += char
				continue
			}
			break
		}

		// Remove spaces if we are not inside quotes
		if char == " " {
			if escapeCount%2 == 1 {
				newtext += char
			}
			continue
		}

		// Remove tabs if we are not inside quotes
		if char == "	" {
			if escapeCount%2 == 1 {
				newtext += char
			}
			continue
		}

		newtext += char
	}

	return newtext
}
