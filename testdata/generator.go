package testdata

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
	"time"
)

const alplhabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alplhabet)

	for i := 0; i < n; i++ {
		c := alplhabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomEmail() string {
	return fmt.Sprintf("%s@example.com", RandomString(6))
}

func TimeFrom(dateStr string) (*time.Time, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func RandomURL() string {
	domain := RandomString(5 + rand.Intn(10)) // 5-14文字のドメイン名
	tld := RandomString(2 + rand.Intn(3))     // 2-4文字のTLD

	return fmt.Sprintf("https://%s.%s", domain, tld)
}

func RandomISBN() string {
	var sb strings.Builder
	for i := 0; i < 13; i++ {
		digit := rand.Intn(10)
		sb.WriteString(fmt.Sprintf("%d", digit))
	}
	return sb.String()
}

func RandomValidPassword() string {
	symbols := "-^$*.@"
	digits := "0123456789"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower := "abcdefghijklmnopqrstuvwxyz"

	// 8-48文字の範囲でランダムな長さを選択
	length := 8 + rand.Intn(41) // 8 + 0~40 = 8~48

	var password strings.Builder

	// 必須文字を1文字ずつ追加
	password.WriteByte(upper[rand.Intn(len(upper))])
	password.WriteByte(lower[rand.Intn(len(lower))])
	password.WriteByte(digits[rand.Intn(len(digits))])
	password.WriteByte(symbols[rand.Intn(len(symbols))])

	// 残りの文字をランダムに生成
	allChars := upper + lower + digits + symbols
	for i := 4; i < length; i++ {
		password.WriteByte(allChars[rand.Intn(len(allChars))])
	}

	// 文字をシャッフル
	passwordBytes := []byte(password.String())
	for i := len(passwordBytes) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		passwordBytes[i], passwordBytes[j] = passwordBytes[j], passwordBytes[i]
	}

	return string(passwordBytes)
}
