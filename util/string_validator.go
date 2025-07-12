package util

import (
	"fmt"
	"regexp"
)

var (
	usernameRegex     = regexp.MustCompile(`^[A-Za-z0-9]{5,30}$`)
	emailRegex        = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	upperCaseRegex    = regexp.MustCompile(`[A-Z]`)
	lowerCaseRegex    = regexp.MustCompile(`[a-z]`)
	digitRegex        = regexp.MustCompile(`[0-9]`)
	symbolRegex       = regexp.MustCompile(`[\-^$*.@]`)
	urlRegex          = regexp.MustCompile(`^https?://[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,}(/.*)?$`)
	isbnRegex         = regexp.MustCompile(`^[0-9]{13}$`)
	sqlInjectionRegex = regexp.MustCompile(`(?i)(select|insert|update|delete|drop|create|alter|exec|union|script|javascript|<script|</script>|--|;|/\*|\*/|'.*'|".*")`)
)

type StringValidator string

func (s StringValidator) ValidateLength(minLength int, maxLength int) error {
	if len(s) < minLength || len(s) > maxLength {
		return fmt.Errorf("length must be between %d and %d", minLength, maxLength)
	}
	return nil
}

func (s StringValidator) validateRegex(re *regexp.Regexp) error {
	if !re.MatchString(string(s)) {
		return fmt.Errorf("'%s' invalid format", s)
	}
	return nil
}

func (s StringValidator) ValidateUsername() error {
	//5文字以上30文字以内 英大文字（A-Z） 英小文字（a-z） 数字（0-9）
	return s.validateRegex(usernameRegex)
}

func (s StringValidator) ValidateEmail() error {
	return s.validateRegex(emailRegex)
}

func (s StringValidator) ValidatePassword() error {
	// 大小英数字記号をそれぞれ1文字以上含む8文字以上48文字以下の文字列
	if s.ValidateLength(8, 48) != nil {
		return fmt.Errorf("password length must be between 8 and 48")
	}
	if !s.hasUpperCase() {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !s.hasLowerCase() {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !s.hasDigit() {
		return fmt.Errorf("password must contain at least one digit")
	}
	if !s.hasSymbol() {
		return fmt.Errorf("password must contain at least one symbol")
	}
	return nil
}

func (s StringValidator) hasUpperCase() bool {
	err := s.validateRegex(upperCaseRegex)
	return err == nil
}

func (s StringValidator) hasLowerCase() bool {
	err := s.validateRegex(lowerCaseRegex)
	return err == nil
}

func (s StringValidator) hasDigit() bool {
	err := s.validateRegex(digitRegex)
	return err == nil
}

func (s StringValidator) hasSymbol() bool {
	// -^$*.@ のいずれかの文字を含む
	err := s.validateRegex(symbolRegex)
	return err == nil
}

func (s StringValidator) ValidateURL() error {
	return s.validateRegex(urlRegex)
}

func (s StringValidator) ValidateISBN() error {
	return s.validateRegex(isbnRegex)
}

func (s StringValidator) ValidateNoSQLInjection() error {
	if sqlInjectionRegex.MatchString(string(s)) {
		return fmt.Errorf("potentially dangerous content detected")
	}
	return nil
}
