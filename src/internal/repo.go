package internal

import (
	"time"

	"github.com/ssibrahimbas/cli-todos/src/entity"
)

type Repo struct {
	todos []*entity.Todo
}

type RepoConfig struct {}

func NewRepo(c *RepoConfig) *Repo {
	return &Repo{
		todos: []*entity.Todo{},
	}
}

func (r *Repo) Add(t *entity.Todo) {
	r.todos = append(r.todos, t)
}

func (r *Repo) Complete(id int) {
	if id >= len(r.todos) || r.todos[id].Done {
		return
	}
	r.todos[id].Done = true
	r.todos[id].UpdatedAt = time.Now()
}

func (r *Repo) Delete(id int) {
	if id >= len(r.todos) {
		return
	}
	r.todos = append(r.todos[:id], r.todos[id+1:]...)
}

func (r *Repo) List() []*entity.Todo {
	return r.todos
}

func (r *Repo) ListNotComplete() []*entity.Todo {
	var todos []*entity.Todo
	for _, t := range r.todos {
		if !t.Done {
			todos = append(todos, t)
		}
	}
	return todos
}

func (r *Repo) ListComplete() []*entity.Todo {
	var todos []*entity.Todo
	for _, t := range r.todos {
		if t.Done {
			todos = append(todos, t)
		}
	}
	return todos
}

func (r *Repo) CountComplete() int {
	var count int
	for _, t := range r.todos {
		if t.Done {
			count++
		}
	}
	return count
}

func (r *Repo) CountNotComplete() int {
	var count int
	for _, t := range r.todos {
		if !t.Done {
			count++
		}
	}
	return count
}