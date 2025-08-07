package main

import (
	"fmt"
	"log"
	"todo/controllers"
	"todo/handler"

	"github.com/gin-gonic/gin"
)

// func healthHandler(w http.ResponseWriter, r *http.Request) {
// 	health := HealthResponse{
// 		Status:  "OK",
// 		Message: "Api handle OK Ji!",
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	json.NewEncoder(w).Encode(health)

// }
// func todosHandler(w http.ResponseWriter, r *http.Request) {
// 	//fmt.Println(r.Method)
// 	switch r.Method {
// 	case "GET":
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(todos)
// 	case "POST":
// 		var newTodo Todo
// 		body, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			http.Error(w, "unable to read from req body", http.StatusBadRequest)
// 			return
// 		}

// 		//fmt.Println(body)
// 		err = json.Unmarshal(body, &newTodo)
// 		if err != nil || newTodo.Task == "" {
// 			http.Error(w, "no inputs or invalid", http.StatusBadRequest)
// 			return
// 		}
// 		newTodo.Id = generateRandomId()

// 		todoMutex.Lock()
// 		todos = append(todos, newTodo)

// 		todoMutex.Unlock()
// 		w.WriteHeader(http.StatusCreated)
// 		json.NewEncoder(w).Encode(newTodo)

// 	default:
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}
// }

// func todosByIdHandler(w http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Path[len("/todos/"):]

// 	todoMutex.Lock()
// 	defer todoMutex.Unlock()

// 	for i, todo := range todos {

// 		if todo.Id == id {
// 			switch r.Method {
// 			case "GET":
// 				w.Header().Set("Content-Type", "application/json")
// 				json.NewEncoder(w).Encode(todo)
// 			case "PUT":
// 				var updatedTodo Todo
// 				body, err := ioutil.ReadAll(r.Body)
// 				if err != nil {
// 					http.Error(w, "Updation by id error", http.StatusBadRequest)
// 					return
// 				}
// 				err = json.Unmarshal(body, &updatedTodo)
// 				if err != nil || updatedTodo.Task == "" {
// 					http.Error(w, "Invalid input", http.StatusBadRequest)
// 					return
// 				}
// 				todos[i].Task = updatedTodo.Task
// 				todos[i].Completed = updatedTodo.Completed

// 				json.NewEncoder(w).Encode(todos[i])

// 			case "DELETE":
// 				deletedTodoId := todo.Id
// 				todos = append(todos[:i], todos[i+1:]...)
// 				w.Header().Set("Content-Type", "application/json")
// 				w.WriteHeader(http.StatusOK)
// 				json.NewEncoder(w).Encode(map[string]string{"message": "Todo with Id" + deletedTodoId + "is successfully deleted"})

// 			default:
// 				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 			}
// 		}

// 	}
// }

func main() {

	router := gin.Default()

	//http.HandleFunc("/health", healthHandler)
	//http.HandleFunc("/todos", todosHandler)
	//http.HandleFunc("/todos/", todosByIdHandler)
	//http.HandleFunc("/todos",todoHandler)

	router.GET("/health", handler.GetHealth)
	controllers.TodosController(router)

	fmt.Println("Server started in port 3000")

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal("serious error", err)
		return
	}

	/*err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("error starting the application", err)
	}*/

}
