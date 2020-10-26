package main

import (
	"storyReadCounter/pkg"
	"storyReadCounter/web"
	"storyReadCounter/web/controller"
)

func main() {
	db := pkg.NewDbClient()
	pkg.DbMigration(db)
	pkg.SeedDB(db)
	cnt := controller.NewController(db)
	ws := web.NewWebServer(cnt)

	ws.InitRoutes()
}
