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
	SPACE = 32
)

func ConvertRNA(value string, w http.ResponseWriter, r *http.Request) string {
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
		case SPACE:
			converted[i] = SPACE
		default:
			http.Redirect(w, r, "/", 308)
		}
	}

	return string(convertedValue)
}

func ConvertDNA(value string, w http.ResponseWriter, r *http.Request) string {
	convertedValue := []rune(strings.ToUpper(value))
	converted = convertedValue
	for i := 0; i < len(convertedValue); i++ {
		switch convertedValue[i] {
		case U:
			converted[i] = A
		case G:
			converted[i] = C
		case A:
			converted[i] = T
		case C:
			converted[i] = G
		case SPACE:
			converted[i] = SPACE
		default:
			http.Redirect(w, r, "/", 308)
		}
	}

	return string(convertedValue)
}

func Protein(mRNA string, w http.ResponseWriter, r *http.Request) string {

	fields := strings.Fields(mRNA)
	proteins := fields

	for i, value := range fields {
		if len(value) < 3 {
			continue
		} else {
			proteins[i] = GetProtein(value, w, r)
		}
	}

	return strings.Join(proteins, " ")
}

func AddSpace(value string) string {
	for i := 3; i < len(value); i += 4 {
		value = value[:i] + " " + value[i:]
	}

	return strings.ToUpper(value)
}