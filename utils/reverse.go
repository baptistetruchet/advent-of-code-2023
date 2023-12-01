package utils

import (
	"strings"

	"github.com/samber/lo"
)

func ReverseString(s string) string {
	return strings.Join(lo.Reverse(strings.Split(s, "")), "")
}
