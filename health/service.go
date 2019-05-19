package health

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// LoopTodos *
func LoopTodos(loop int) []Todo {
	var todos []Todo
	var todo Todo
	for i := 1; i <= loop; i++ {
		url := "https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(int(i))
		res, _ := http.Get(url)
		temp, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(temp, &todo)
		todos = append(todos, todo)
	}
	return todos
}
