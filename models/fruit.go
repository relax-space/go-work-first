package models

import "fmt"

func Hello(mod string) string {
	return fmt.Sprintf("hello %v => models", mod)
}
