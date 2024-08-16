package hello_world

const (
	spanish = "Spanish"
	france  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	franceHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	greetingPrefix := greetingPrefix(language)
	return greetingPrefix + name
}

func greetingPrefix(language string) (greetingPrefix string) {
	switch language {
	case spanish:
		greetingPrefix = spanishHelloPrefix
	case france:
		greetingPrefix = franceHelloPrefix
	default:
		greetingPrefix = englishHelloPrefix
	}
	return
}
