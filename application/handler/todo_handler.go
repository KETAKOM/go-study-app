package handler

type todo struct {
	title  string
	detail string
}

func NewTodo() *todo {
	return &todo{"", ""}
}

func (t *todo) Create(ti, da string) *todo {
	t.title = ti
	t.detail = da
	return t
}

func (t *todo) Show() *todo {
	return t
}
