package main

func main() {
	// 先创建一只火鸡
	var turkey *WildTurkey = &WildTurkey{}

	// 再将火鸡包装进一个火鸡适配器中，使它看起来像是一只鸭子。
	var turkeyAdapter Duck = NewTurkeyAdapter(turkey)

	// 测试鸭子方法，传入一只假装是鸭子的火鸡
	testDuck(turkeyAdapter)
}

func testDuck(duck Duck) {
	duck.Quack()
	duck.Fly()
}
