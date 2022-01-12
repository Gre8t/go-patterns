package main

type Athlete struct{}

func (a *Athlete) Train() {
	println("Training")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    *func()
}

func Swim() {
	println("Swimming!")
}

var localSwim = Swim

type Animal struct{}

func (r *Animal) Eat() {
	println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}

type Swimmer interface {
	Swim()
}
type Trainer interface {
	Train()
}
type SwimmerImplementor struct{}

func (s *SwimmerImplementor) Swim() {
	println("Swimming!")
}

type CompositeSwimmerB struct {
	Athlete
	Swimmer
}

func main() {

	var swimmer = CompositeSwimmerB{
		Athlete: Athlete{},
		Swimmer: &SwimmerImplementor{},
	}

	swimmer.Train()
	(swimmer.Swim)()
}
