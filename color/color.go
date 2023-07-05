package color

import "strings"

// Minecraft color and format codes are prefixed by this character ("§")
const COLOR_CODE_PREFIX string = "§"

// map Minecraft color and format codes to ANSI escape codes
var COLOR_CODES map[string]string = map[string]string{
	"0": "\033[0;30m", // Black
	"1": "\033[0;34m", // Dark Blue
	"2": "\033[0;32m", // Dark Green
	"3": "\033[0;36m", // Cyan
	"4": "\033[0;31m", // Dark Red
	"5": "\033[0;35m", // Purple
	"6": "\033[0;33m", // Gold
	"7": "\033[0;37m", // Light Gray
	"8": "\033[1;30m", // Gray
	"9": "\033[1;34m", // Blue
	"a": "\033[1;32m", // Green
	"b": "\033[1;36m", // Aqua
	"c": "\033[1;31m", // Red
	"d": "\033[1;35m", // Pink
	"e": "\033[1;33m", // Yellow
	"f": "\033[1;37m", // White

	"k": "\033[5m", // Obfuscated
	"l": "\033[1m", // Bold
	"m": "\033[9m", // Strikethrough
	"n": "\033[4m", // Underline
	"o": "\033[3m", // Italic
	"r": "\033[0m", // Reset
}

// replaces all Minecraft color and format codes with ANSI escape codes
func ParseColorCodes(str string, colors bool) string {
	if colors {
		for code, ansi := range COLOR_CODES {
			str = strings.ReplaceAll(str, COLOR_CODE_PREFIX+code, ansi)
		}
		str += COLOR_CODES["r"]
	} else {
		for code := range COLOR_CODES {
			str = strings.ReplaceAll(str, COLOR_CODE_PREFIX+code, "")
		}
	}
	return str
}

// translate a Minecraft color or format code to an ANSI code
func TranslateColorCode(colorCode string) string {
	return COLOR_CODES[colorCode]
}
