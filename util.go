package main

import (
	"crypto/rand"
	"encoding/hex"
	mrand "math/rand"
	"os"
	"strings"
)

func randKey(l int) []byte {
	buf := make([]byte, l)
	mrand.Read(buf)
	return buf
}

func generateKeys(c, min, max int) [][]byte {
	keys := make([][]byte, 0, c)
	for len(keys) < c {
		if max-min == 0 {
			keys = append(keys, randKey(min))
		} else {
			kl := mrand.Intn(max-min) + min
			keys = append(keys, randKey(kl))
		}
	}
	return keys
}

func randVal() string {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(buf)
}

func padStart(s string, l int) string {
	if len(s) >= l {
		return s
	}
	return strings.Repeat("0", l-len(s)) + s
}

func shuffle(a [][]byte) {
	for i := len(a) - 1; i > 0; i-- {
		j := mrand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func DirSize(path string) (int64, error) {
	var total int64
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}
	for _, e := range entries {
		if e.Type().IsDir() {
			n, err := DirSize(path + "/" + e.Name())
			if err != nil {
				return 0, err
			}
			total += n
		}
		if e.Type().IsRegular() {
			info, err := e.Info()
			if err != nil {
				return 0, nil
			}
			total += info.Size()
		}
	}
	return total, nil
}
