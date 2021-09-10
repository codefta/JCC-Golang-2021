package tech

import (
	"fmt"
	"strings"
)

type Phone struct {
	Name, Brand string
	Year int
	Colors []string
}

type PhoneSpec interface {
	ShowPhoneSpec() string
}

func (p Phone) ShowPhoneSpec() string {
	return fmt.Sprintf("name : %v\nbrand : %v\nyear : %d\ncolors : %v", p.Name, p.Brand, p.Year, strings.TrimRight(strings.Join(p.Colors, ", "), ", "))
}