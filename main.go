package main

import (
	"github.com/kataras/go-template/html"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/kataras/iris"
)

type mypage struct {
	Title   string
	Message string
}

func main() {

	iris.UseTemplate(html.New(html.Config{
		Layout: "layout.html",
	})).Directory("./templates", ".html") // the .Directory() is optional also, defaults to ./templates, .html
	// Note for html: this is the default iris' templaet engine, if zero engines added, then the template/html will be used automatically
	// These lines are here to show you how you can change its default configuration

	iris.Get("/", func(ctx *iris.Context) {
		db, _:= sql.Open("mysql", "/test")
		
		rows, _ := db.Query("SELECT username FROM users")
		var username string
		for rows.Next() {
			err, _ := rows.QueryRow(1).Scan(&username)
			fmt.Println(username)
		}
		
		
		ctx.Render("mypage.html", mypage{"My Page title", "Hello world!"}, iris.RenderOptions{"gzip": true})
		// Note that: you can pass "layout" : "otherLayout.html" to bypass the config's Layout property or iris.NoLayout to disable layout on this render action.
		// RenderOptions is an optional parameter
	})

	iris.Listen(":8080")
}