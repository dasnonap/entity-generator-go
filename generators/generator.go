package generators

import (
	"fmt"
	"math/rand"
	"regexp/syntax"
	"strings"
)

const runeRangeEnd = 0x10ffff
const printableChars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ \t\n\r"
var printableCharsNoNL = printableChars[:len(printableChars)-2]

func Generate(tokens *syntax.Regexp) (string) {
	operator := tokens.Op

	switch operator {
		case syntax.OpCharClass:
			sum := 0
			for i := 0; i < len(tokens.Rune); i += 2 {
				sum += int(tokens.Rune[i+1] - tokens.Rune[i]) + 1

				if tokens.Rune[i+1] == runeRangeEnd {
					sum = -1
					break
				}
			}
			
			if sum == -1 {
				fmt.Printf("Sum is zero");
			}
			position := rand.Intn(int(sum))

			var ru rune;
			sum = 0;

			for i := 0; i < len(tokens.Rune); i += 2 {
				gap := int(tokens.Rune[i+1] - tokens.Rune[i]) + 1

				if gap+sum > position {
					ru = tokens.Rune[i] + rune(position-sum)
					break
				}

				sum += gap
			}
			return string(ru)
		case syntax.OpAnyCharNotNL:
		case syntax.OpLiteral:
			res := ""
			for _, value := range tokens.Rune {
				if strings.Contains(printableChars, string(value)) {
					res += string(value)
				}
			}
			return res
		case syntax.OpConcat:
			res := ""
			for _, regex := range tokens.Sub {
				res += Generate(regex);		
			}

			return res
		case syntax.OpCapture:
			return Generate(tokens.Sub[0])
		case syntax.OpAlternate:
			index := rand.Intn(len(tokens.Sub))

			return Generate(tokens.Sub[index])
		case syntax.OpRepeat:
			min := tokens.Min
			max := tokens.Max

			if min < 0  {
				min *= -1
			}

			if max < 0 {
				max *= -1
			}

			size := rand.Intn(max-min) + min
			res := ""

			for i := 0; i < size; i++ {
				for _, value := range tokens.Sub {
					res += Generate(value)
				}
			}

			return res
		default:
			panic("No syntax found for class" + tokens.Op.String())
	}

	// return tokens, error
	return ""
}