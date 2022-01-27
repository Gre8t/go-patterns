package main

import "fmt"

type Command interface {
	Execute()
}
type ConsoleOutput struct {
	message string
	msg string
}

func (c *ConsoleOutput) Execute(){
	fmt.Println(c.message)
}

type CommandQueue struct{
	queue []Command
}
func (p *CommandQueue) AddCommand(c Command){
	p.queue = append(p.queue, c)
	if len(p.queue) == 3 {
		for _, command := range p.queue {
			command.Execute()
		}
		//this is mad ohhh, just experienced the thrill of your first refactor while learning from someone else code
		//ok, this feels euphoric to me. This is why it's always important to have a reviewer for your book.
		p.queue = make([]Command, 0, 3)
	}
}
func CreateCommand(str string) Command{
	co := ConsoleOutput{message: str,}
	return &co
}

func main() {
	queue := CommandQueue{}

	queue.AddCommand(CreateCommand("First message"))
	queue.AddCommand(CreateCommand("Second message"))
	queue.AddCommand(CreateCommand("Third message"))
	queue.AddCommand(CreateCommand("Fourth message"))
	queue.AddCommand(CreateCommand("Fifth message"))
	queue.AddCommand(CreateCommand("sixth message"))
}
