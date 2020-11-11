package strutil

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/fatih/camelcase"

	"github.com/dinalt/toolbox/stringlist"
)

// TruncRight cuts right part of string "s" to make it length of "length"
func TruncRight(s string, length int) string {
	runes := []rune(s)
	if len(runes) > length {
		return string(runes[:length+1])
	}
	return string(runes)
}

// ToSnakeCase converts camelCase string to snake_case string
func ToSnakeCase(s string) string {
	sl := stringlist.New(0)
	for _, r := range camelcase.Split(s) {
		sl.Add(strings.ToLower(r))
	}
	return sl.Join("_")
}

// ToKVPairs separate string to keys and values using pairDelim and kvDelim
func ToKVPairs(s, pairDelim, kvDelim string) (map[string]string, error) {
	params := make(map[string]string)
	for _, p := range strings.Split(strings.ToLower(s), pairDelim) {
		kv := strings.Split(p, kvDelim)
		if len(kv) > 2 {
			return params, fmt.Errorf("ToKVPairs(): String has wrong format: \"%s\"", s)
		}
		if len(kv) == 2 {
			params[kv[0]] = kv[1]
		} else {
			params[kv[0]] = ""
		}
	}

	return params, nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    //  6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 //  All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   //  # of letter indices fitting in 63 bits
)

// RandStringBytes returns string of random ASCII chars. It uses math/rand
// instead of crypto/rand, so if you need strong RNG, this is not the case.
// Taken from this StackOverflow answer: https:// stackoverflow.com/a/31832326
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; { // nolint:gosec
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax // nolint:gosec
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
