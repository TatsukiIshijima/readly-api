package util

import (
	"fmt"
	"regexp"
)

type StringValidator string

func (s StringValidator) ValidateLength(minLength int, maxLength int) error {
	if len(s) < minLength || len(s) > maxLength {
		return fmt.Errorf("length must be between %d and %d", minLength, maxLength)
	}
	return nil
}

func (s StringValidator) validateRegex(pattern string) error {
	re := regexp.MustCompile(pattern)
	if !re.MatchString(string(s)) {
		return fmt.Errorf("'%s' invalid format", s)
	}
	return nil
}

func (s StringValidator) ValidateUsername() error {
	//5文字以上30文字以内 英大文字（A-Z） 英小文字（a-z） 数字（0-9）
	return s.validateRegex(`^[A-Za-z0-9]{5,30}$`)
}

func (s StringValidator) ValidateEmail() error {
	return s.validateRegex(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
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
	err := s.validateRegex(`[A-Z]`)
	return err == nil
}

func (s StringValidator) hasLowerCase() bool {
	err := s.validateRegex(`[a-z]`)
	return err == nil
}

func (s StringValidator) hasDigit() bool {
	err := s.validateRegex(`[0-9]`)
	return err == nil
}

func (s StringValidator) hasSymbol() bool {
	// -^$*.@ のいずれかの文字を含む
	err := s.validateRegex(`[\-^$*.@]`)
	return err == nil
}
