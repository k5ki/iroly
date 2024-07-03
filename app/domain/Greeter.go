package domain

type Hello struct {
	Message string `json:"message"`
}

type Greeter struct{}

func NewGreeter() *Greeter {
	return &Greeter{}
}

func (g *Greeter) Hello() *Hello {
	return &Hello{Message: "Hello, World!"}
}
