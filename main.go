package main

func main() {
	//app, err := initApp()
	//must(err)
	//

	c, _ := ProviderCCCCCC()
	c.View()
	c.B().Show()
	c.A().Show()
	c.View()

	//must(app.Start())

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
