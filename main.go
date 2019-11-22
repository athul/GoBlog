package main

import (
	"io/ioutil"
	"github.com/gin-gonic/gin"
	//"html/template"
	"fmt"
	"net/http"
	"log"
)
func main(){
	r:=gin.Default()
	r.Use(gin.Logger())
	r.Delims("{{","}}")
	r.LoadHTMLGlob("./templates/*.tmpl.html")

	r.GET("/",func (c *gin.Context){
		var posts []string
		files,err:=ioutil.ReadDir("./markdown/")
		if err!=nil{
			log.Fatal(err)
		}
		for _,file:=range files{
			fmt.Println(file.Name())
			posts=append(posts,file.Name())
		}
		c.HTML(http.StatusOK,"index.tmpl.html",gin.H{
			"posts":posts,
		})
	})
	r.Run()

}