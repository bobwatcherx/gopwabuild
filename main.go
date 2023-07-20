package main

import (
	"log"
	"github.com/maxence-charriere/go-app/v9/pkg/app"

)

type hello struct{
	app.Compo
	name string
}

func(c *hello) Render() app.UI{
	c.name = "youtube kids"

	return app.Div().
	Style("background-color","red").
	Style("padding","10px").
	Body(
		app.H1().Body(
			app.Text("Hai , "),
			app.If(c.name != "",
				app.Text(c.name),
			).Else(
				app.Text("Say yes"),
			),

		),

	)

}


// NOW CREATE SERVER 
func main(){
	app.Route("/",&hello{})
	app.RunWhenOnBrowser()

	err :=app.GenerateStaticWebsite(".",&app.Handler{
		Name: "Hello app",
		Description : "this app pwa hello",
		Resources : app.GitHubPages("gopwabuild"),

	})
	 if err != nil{
	 	log.Fatal(err)
	 }
}