package factory_method

import "fmt"

type Fruit interface {
	Eat()
}

type Orange struct {
	name string
}

func NewOrange(name string) Fruit {
	return &Orange{
		name: name,
	}
}

func (o *Orange) Eat() {
	fmt.Printf("i am orange: %s, i am about to be eaten... \n", o.name)
}

type Strawberry struct {
	name string
}

func NewStrawberry(name string) Fruit {
	return &Strawberry{
		name: name,
	}
}

func (s *Strawberry) Eat() {
	fmt.Printf("i am strawberry: %s, i am about to be eaten... \n", s.name)
}

type Cherry struct {
	name string
}

func NewCherry(name string) Fruit {
	return &Cherry{
		name: name,
	}
}

func (c *Cherry) Eat() {
	fmt.Printf("i am cherry: %s, i am about to be eaten... \n", c.name)
}

type FruitFactory interface {
	CreateFruit(name string) Fruit
}

type OrangeFactory struct {
}

func (of *OrangeFactory) CreateFruit(name string) Fruit {
	return NewOrange(name)
}

func NewOrangeFactory() FruitFactory {
	return &OrangeFactory{}
}

type StrawberryFactory struct {
}

func (sf *StrawberryFactory) CreateFruit(name string) Fruit {
	return NewStrawberry(name)
}

func NewStrawberryFactory() FruitFactory {
	return &StrawberryFactory{}
}

type CherryFactory struct {
}

func (cf *CherryFactory) CreateFruit(name string) Fruit {
	return NewCherry(name)
}

func NewCherryFactory() FruitFactory {
	return &CherryFactory{}
}
