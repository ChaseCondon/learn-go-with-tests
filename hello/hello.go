package main

import "fmt"

var greetings = map[string]string{
	"English": "Hello, ",
	"Spanish": "Hola, ",
	"French":  "Bonjour, ",
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "" {
		language = "English"
	}

	greeting, ok := greetings[language]
	if ok {
		return fmt.Sprintf("%s%s!", greeting, name)
	} else {
		panic(fmt.Sprintf("%s is not a supported language", language))
	}
}

func main() {
	fmt.Println(Hello("Chase", ""))
}
