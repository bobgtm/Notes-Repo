package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name, language string) string {
	greeting := ""
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		greeting = spanishHelloPrefix + name
	}
	if language == "English" {
		greeting = englishHelloPrefix + name
	}
	return greeting
}

func main() {
	fmt.Println(Hello("Bobby", "Spanish"))
}
