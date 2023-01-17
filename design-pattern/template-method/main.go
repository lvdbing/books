package main

func main() {
	var tea *Tea = &Tea{}
	tea.beverage = tea
	tea.PrepareRecipe()

	var coffee *Coffee = &Coffee{}
	coffee.beverage = coffee
	coffee.PrepareRecipe()
}
