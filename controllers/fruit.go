package controllers

import (
	"fmt"
	"models"
)

func Hello(mod string) string {
	return models.Hello(fmt.Sprintf("%v => controllers", mod))
}
