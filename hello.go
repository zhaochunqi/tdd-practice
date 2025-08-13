package main

import "fmt"

const spanish = "spanish"
const chinese = "chinese"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const chineseHelloPrefix = "你好, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
