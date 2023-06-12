package factory_method

import "testing"

func Test_factory(t *testing.T) {
	//吃个橘子
	orangeFactory := NewOrangeFactory()
	orange := orangeFactory.CreateFruit("1号橘子")
	orange.Eat()

	//吃个樱桃
	cherryFactory := NewCherryFactory()
	cherry := cherryFactory.CreateFruit("2号樱桃")
	cherry.Eat()

	//需要加入新的水果时 只需要完成Fruit接口和创造该水果的工厂FruitFactory接口 就可以先创造对应的工厂 然后由工厂创造对应的水果
}
