package main

import (
	"fmt"
	"reflect"
)

func main() {
	// =======
	// ======= reflect.Type 反射类型对象
	// =======
	// ======= 类型对象 reflect.StructField 和 reflect.Method

	typeOfHero := reflect.TypeOf(hero{})
	fmt.Printf("Heor's type is %s, kind is %s\n",
		typeOfHero, typeOfHero.Kind())

	typeOfPtrHero := reflect.TypeOf(&hero{})
	fmt.Printf("*Hero's type is %s, kind is %s\n",
		typeOfPtrHero, typeOfPtrHero.Kind())
	fmt.Printf("*Hero's Elem type is %s, kind is %s\n",
		typeOfPtrHero.Elem(), typeOfPtrHero.Elem().Kind())

	// 通过 NumField 获取结构体字段的数量
	for i := 0; i < typeOfHero.NumField(); i++ {
		fmt.Printf("field's name is %s, type is %s, kind is %s, tag is %s\n",
			typeOfHero.Field(i).Name,
			typeOfHero.Field(i).Type,
			typeOfHero.Field(i).Type.Kind(),
			typeOfHero.Field(i).Tag)
	}
	// 获取名称为 Name 的成员字段类型对象
	nameField, _ := typeOfHero.FieldByName("Name")
	fmt.Printf("FieldByName got field's name is %s, type is %s, kind is %s, tag is %s\n",
		nameField.Name,
		nameField.Type,
		nameField.Type.Kind(),
		nameField.Tag)

	// 声明一个 person 接口, 并用 hero 作为接收器
	var p person = &hero{}
	typeOfPerson := reflect.TypeOf(p) // 获取接口person的类型对象
	for i := 0; i < typeOfPerson.NumMethod(); i++ {
		fmt.Printf("method is %s, type is %s, kind is %s\n",
			typeOfPerson.Method(i).Name,
			typeOfPerson.Method(i).Type,
			typeOfPerson.Method(i).Type.Kind())
	}
	m, _ := typeOfPerson.MethodByName("Run")
	fmt.Printf("method is %s, type is %s, kind is %s\n",
		m.Name, m.Type, m.Type.Kind())

	// reflect.Type 调用方法, 需要显式将方法的接收器放在 in []Value 参数列表的第一位
	var apple person = &hero{Name: "小苹果"}
	typeOfApple := reflect.TypeOf(apple)
	runMd, _ := typeOfApple.MethodByName("Run")
	// 将 person 接收器放在参数的第一位
	callResult := runMd.Func.Call([]reflect.Value{reflect.ValueOf(apple)})
	fmt.Printf("result of run method is: %s\n", callResult[0])

	fmt.Printf("\n\n\n")

	// =======
	// ======= reflect.Value 反射值对象
	// =======

	name := "小明"
	valueOfName := reflect.ValueOf(name)
	fmt.Println(valueOfName.Interface())

	// 可以用 reflect.New 创建一个相同类型的新变量
	heroValue := reflect.New(typeOfHero)
	fmt.Printf("Hero's type is %s, kind is %s\n", heroValue.Type(), heroValue.Kind())

	// Value.Set 可以修改变量
	valueOfPtrName := reflect.ValueOf(&name)
	valueOfPtrName.Elem().Set(reflect.ValueOf("小红"))
	fmt.Println(name)

	// Value.CanAddr 来判断一个变量的 Value 是否可寻址
	fmt.Printf("name can be address: %t\n", valueOfName.CanAddr())
	fmt.Printf("&name can be address: %t\n", valueOfPtrName.CanAddr())
	fmt.Printf("&name's Elem can be address: %t\n", valueOfPtrName.Elem().CanAddr())

	// Value.CanSet 可以判断变量 Value 是否可以设置
	h := &hero{Name: "小白", pri: "这是隐私不能修改"}
	vOfHero := reflect.ValueOf(h).Elem()
	vOfName := vOfHero.FieldByName("Name")
	if vOfName.CanSet() {
		vOfName.Set(reflect.ValueOf("校长"))
	}
	vOfPri := vOfHero.FieldByName("pri")
	if vOfPri.CanSet() {
		vOfPri.Set(reflect.ValueOf("隐私被修改了"))
	}
	fmt.Printf("hero name is %s, hero pri is %s\n", h.Name, h.pri)

	// Value.Call 可以调用函数
	var monkey person = &hero{Name: "小猴子", Speed: 100}
	vOfMonkey := reflect.ValueOf(monkey)
	// 获取 SayHello 方法
	sayHelloMethod := vOfMonkey.MethodByName("SayHello")
	// 构建调用参数并通过 Call 调用方法
	r := sayHelloMethod.Call([]reflect.Value{reflect.ValueOf("小老虎")})
	fmt.Printf("result of SayHello method is: %v\n", r)
	runMethod := vOfMonkey.MethodByName("Run")
	// 通过 Call 调用方法并获取结果
	result := runMethod.Call([]reflect.Value{})
	fmt.Printf("result of run method is: %s\n", result[0])
}
