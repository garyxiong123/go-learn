package main

func main() {

	cat := &Cat{
		Animal: &Animal{
			Name: "cat",
		},
		Animal1: &Animal1{
			Name: "cat",
		},
	}
	cat.Eat()

	//Sum()
}
