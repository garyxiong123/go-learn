package tclass

func main1() {

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
