package helpers

import (
	"net/http"
	"strings"
)

var converted []rune

const (
	A = 65
	C = 67
	T = 84
	U = 85
	G = 71
)

func Convert(value string, w http.ResponseWriter, r *http.Request) string {
	convertedValue := []rune(strings.ToUpper(value))
	converted = convertedValue
	for i := 0; i < len(convertedValue); i++ {
		switch convertedValue[i] {
		case A:
			converted[i] = U
		case C:
			converted[i] = G
		case T:
			converted[i] = A
		case G:
			converted[i] = C
		default:
			http.Redirect(w, r, "/", 308)
		}
	}

	return string(convertedValue)
}

func Protein(mRNA string, w http.ResponseWriter, r *http.Request) {

}