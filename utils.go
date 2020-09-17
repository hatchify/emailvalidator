package emailvalidator

// Local part characters can be any of the following:
//	- uppercase and lowercase Latin letters A to Z and a to z;
//	- digits 0 to 9;
//	- .

func isValidLocalPartChar(r rune) (ok bool) {
	switch {
	case isLowercaseAlpha(r):
	case isUppercaseAlpha(r):
	case isDigit(r):
	case isAllowedSpecialChar(r):

	default:
		return false
	}

	return true
}

func isLowercaseAlpha(r rune) (ok bool) {
	switch r {
	case 'a':
	case 'b':
	case 'c':
	case 'd':
	case 'e':
	case 'f':
	case 'g':
	case 'h':
	case 'i':
	case 'j':
	case 'k':
	case 'l':
	case 'm':
	case 'n':
	case 'o':
	case 'p':
	case 'q':
	case 'r':
	case 's':
	case 't':
	case 'u':
	case 'v':
	case 'w':
	case 'x':
	case 'y':
	case 'z':

	default:
		return false
	}

	return true
}

func isUppercaseAlpha(r rune) (ok bool) {
	switch r {
	case 'A':
	case 'B':
	case 'C':
	case 'D':
	case 'E':
	case 'F':
	case 'G':
	case 'H':
	case 'I':
	case 'J':
	case 'K':
	case 'L':
	case 'M':
	case 'N':
	case 'O':
	case 'P':
	case 'Q':
	case 'R':
	case 'S':
	case 'T':
	case 'U':
	case 'V':
	case 'W':
	case 'X':
	case 'Y':
	case 'Z':

	default:
		return false
	}

	return true
}

func isDigit(r rune) (ok bool) {
	switch r {
	case '0':
	case '1':
	case '2':
	case '3':
	case '4':
	case '5':
	case '6':
	case '7':
	case '8':
	case '9':

	default:
		return false
	}

	return true
}

func isAllowedSpecialChar(r rune) (ok bool) {
	switch r {
	case '.':
	case '_':

	default:
		return false
	}

	return true
}
