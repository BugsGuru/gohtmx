package model

var Db = make(map[string]bool, 10)

type Todo struct {
	Id   uint64 `json:"id"`
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}

func CreateTodo(todo string) error {
	if todo == "" {
		return nil
	}
	Db[todo] = false
	return nil
}

func GetAllTodos() ([]Todo, error) {
	r := make([]Todo, 0, len(Db))
	for k, v := range Db {
		r = append(r, Todo{Todo: k, Done: v})
	}
	return r, nil
}

func GetTodo(todo string) (Todo, error) {
	return Todo{Todo: todo, Done: Db[todo]}, nil
}

func MarkDone(todo string) error {
	Db[todo] = !Db[todo]
	return nil
}

func Delete(todo string) error {
	delete(Db, todo)
	return nil
}
