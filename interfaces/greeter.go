package interfaces

import "fmt"

type Greeter interface {
	Greet(name string) string
}

type EnglishGreeter struct{}

func (e EnglishGreeter) Greet(name string) string {
	return fmt.Sprintf("Hello, %s welcome", name)
}

type HindiGreeter struct{}

func (h HindiGreeter) Greet(name string) string {
	return fmt.Sprintf("Namaste, %s aapka swagat hai", name)
}

func sayHello(g Greeter, name string) {
	fmt.Println(g.Greet(name))
}

func Greeting() {
	eng := EnglishGreeter{}
	hin := HindiGreeter{}

	sayHello(eng, "Biswas")
	sayHello(hin, "Biswas")
}
