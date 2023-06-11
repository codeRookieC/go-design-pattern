package simple_factory

import "testing"

func Test_factory(t *testing.T) {
	//构造工厂
	factory := NewFruitFactory()
	//吃个橘子
	orange, _ := factory.CreateFruit("orange")
	orange.Eat()
	//吃个草莓
	strawberry, _ := factory.CreateFruit("strawberry")
	strawberry.Eat()
	//吃个西瓜 但是没有对应的构造方法 直接报错 捕获错误
	watermelon, err := factory.CreateFruit("watermelon")
	if err != nil {
		t.Error(err)
		return
	}
	watermelon.Eat()
}
