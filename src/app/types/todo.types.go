package types

type TodoResponse struct {
	ID        uint   `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

type CreateDTO struct {
	Task string `json:"task" validate:"required,min=3,max=150"`
}

type TodoCreateResponse struct {
	Todo *TodoResponse `json:"todo"`
}

type TodosResponse struct {
	Todos *[]TodoResponse `json:"todos"`
}

type CheckTodoDTO struct {
	Completed bool `json:"completed"`
}
