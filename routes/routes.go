package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"htmx/model"
	"log"
	"net/http"
)

func sendTodos(w http.ResponseWriter) {
	todos, err := model.GetAllTodos()
	if err != nil {
		fmt.Println("Could not get all todos from db", err)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.gohtml"))

	err = tmpl.ExecuteTemplate(w, "Todos", todos)
	if err != nil {
		fmt.Println("Could not execute template", err)
	}
}

func index(c *gin.Context) {
	todos, err := model.GetAllTodos()
	if err != nil {
		fmt.Println("Could not get all todos from db", err)
	}

	tmpl := template.Must(template.ParseFiles("templates/index.gohtml"))

	err = tmpl.Execute(c.Writer, todos)
	if err != nil {
		fmt.Println("Could not execute template", err)
	}
}

func markTodo(c *gin.Context) {
	err := model.MarkDone(c.Param("id"))
	if err != nil {
		fmt.Println("Could not update todo", err)
	}
	sendTodos(c.Writer)
}

func createTodo(c *gin.Context) {
	r := c.Request
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form", err)
	}

	err = model.CreateTodo(r.FormValue("todo"))
	if err != nil {
		fmt.Println("Could not create todo", err)
	}

	sendTodos(c.Writer)
}

func deleteTodo(c *gin.Context) {
	err := model.Delete(c.Param("id"))
	if err != nil {
		fmt.Println("Could not delete", err)
	}

	sendTodos(c.Writer)

}

func SetupAndRun() {
	r := gin.Default()

	r.GET("/", index)
	r.PUT("/todo/:id", markTodo)
	r.DELETE("/todo/:id", deleteTodo)
	r.POST("/create", createTodo)

	log.Fatal(r.Run(":8000"))
}
