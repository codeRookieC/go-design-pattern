package simple_factory

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Fruit Fruit 是一个抽象的 interface，水果的共性是都可以食用水果，这里我们不纠结主被动的语义，赋以 Fruit 一个 Eat 方法
type Fruit interface {
	Eat()
}

//有三个具体的水果实现类，橘子 Orange、草莓 Strawberry、樱桃 cherry，分别实现了 Fruit 对应的 Eat 方法

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

//有一个具体的水果工厂类 FruitFactory，专门用于水果的生产工作，对应的生产方法为CreateFruit 方法，可以按照用户指定的水果类型，生产出对应的水果
//利用工厂生产三类时存在的公共切面，进行随机数的取值，用来给生产出来的水果命名
//根据使用方传入的水果类型 typ，调用对应水果类型的构造器方法，并将生产出来的水果进行返回
//如果使用方法传入的水果类型 typ 非法，则对外抛出错误

/*type FruitFactory struct {
}

func NewFruitFactory() *FruitFactory {
	return &FruitFactory{}
}

// CreateFruit 根据传入的参数创建对应的Fruit
func (f *FruitFactory) CreateFruit(typ string) (Fruit, error) {
	src := rand.NewSource(time.Now().UnixNano())
	random := rand.New(src)
	name := strconv.Itoa(random.Int())

	switch typ {
	case "orange":
		return NewOrange(name), nil
	case "strawberry":
		return NewStrawberry(name), nil
	case "cherry":
		return NewCherry(name), nil
	default:
		return nil, fmt.Errorf("fruit typ: %s is not supported yet", typ)
	}
}*/

//当需要支持的水果类型的typ数量提升的时候 会产生方法圈复杂度过高的问题 改进主要针对switch-case结构
//方法圈的复杂度过高：方法的控制流程过于复杂，包含了大量的条件分支和循环结构，导致方法难以理解、测试和维护
//这里改进 在FruitFactory中内置一个map; key -> typ, value -> func(name string) Fruit -> NewXXX

type fruitCreator func(name string) Fruit

type FruitFactory struct {
	creators map[string]fruitCreator
}

func NewFruitFactory() *FruitFactory {
	return &FruitFactory{
		creators: map[string]fruitCreator{
			"orange":     NewOrange,
			"strawberry": NewStrawberry,
			"cherry":     NewCherry,
		},
	}
}

func (f *FruitFactory) CreateFruit(typ string) (Fruit, error) {
	fruitCreator, ok := f.creators[typ]
	if !ok {
		return nil, fmt.Errorf("fruit typ: %s is not supported yet \n", typ)
	}
	src := rand.NewSource(time.Now().UnixNano())
	random := rand.New(src)
	name := strconv.Itoa(random.Int())
	return fruitCreator(name), nil
}
