package main

import (
	api2 "github.com/Inf85/go-project-demo/api"
	"github.com/Inf85/go-project-demo/database"
	routes2 "github.com/Inf85/go-project-demo/routes"
	"github.com/urfave/negroni"
)

func main() {
	db := database.NewPostgresDB("host=localhost user=demo_user password=123456 dbname=demo_dev port=54323 sslmode=disable")
	api := api2.NewAPI(db.DB)
	routes := routes2.NewRoutes(api)
	n := negroni.Classic()
	n.UseHandler(routes)
	n.Run(":3000")
}
