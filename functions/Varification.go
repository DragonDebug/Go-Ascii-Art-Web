package functions

import (
	"fmt"
	"regexp"
)

// Varifying if the user option value is valid and return the file name of the
func VarifyOption(userOption string) (string, error) {
	reOption := regexp.MustCompile(`(?i)--output=([^\\/:*?"<>|]{1,244}(\.[a-zA-Z0-9]{1,10}))`)
	reExt := regexp.MustCompile(`(?i)\.txt`)
	reFile := regexp.MustCompile(`(?i)(standard|shadow|thinkertoy)\.txt`)

	// Checking if the option flag is in the correct format
	if !reOption.MatchString(userOption) {
		return "", fmt.Errorf("invalid option: '%s'", userOption)
	}

	submatches := reOption.FindStringSubmatch(userOption)

	// Checking if the file extention is .txt
	extention := submatches[2]
	if !reExt.MatchString(extention) {
		return "", fmt.Errorf("invalid file extention: '%s'", extention)
	}

	// Checking the name of the file
	fileName := submatches[1]
	if reFile.MatchString(fileName) {
		return "", fmt.Errorf("invalid file name: '%s'", fileName)
	}
	return fileName, nil
}

// Varifying if the banner value is valid
func VarifyBanner(userBanner string) (string, error) {
	rebanner := regexp.MustCompile(`(?i)(standard|shadow|thinkertoy)(\.txt)?`)

	if !rebanner.MatchString(userBanner) {
		return "", fmt.Errorf("invalid banner: '%s'. Available banners are: standard, shadow and thinkertoy", userBanner)
	}

	submatches := rebanner.FindStringSubmatch(userBanner)
	formatedbanner := submatches[1] + ".txt"
	return formatedbanner, nil
}

// Checking if the user input consists of only ASCII characters by checking if the key exists inside of the map
func IsAsciiInput(word string, asciiArtMap map[rune][]string) bool {
	for _, char := range word {
		_, exists := asciiArtMap[char]
		if exists {
			continue
		} else {
			return false
		}
	}
	return true
}
