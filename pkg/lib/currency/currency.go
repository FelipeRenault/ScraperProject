package currency

import (
	"regexp"
	"strconv"

	"github.com/leekchan/accounting"
)

func ParsePrice(price string) (int, error) {
	price = accounting.UnformatNumber(price, 2, "BRL")

	r := regexp.MustCompile(`\.`) // Remove decimal separator
	price = r.ReplaceAllString(price, "${1}")

	return strconv.Atoi(price)
}
