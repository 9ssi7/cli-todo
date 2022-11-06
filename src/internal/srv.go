package internal

import (
	"github.com/ssibrahimbas/cli-todos/src/entity"
)

type Srv struct {
	r *Repo
}

type SrvConfig struct {
	Repo *Repo
}

func NewSrv(c *SrvConfig) *Srv {
	return &Srv{
		r: c.Repo,
	}
}

func (s *Srv) Add(t *entity.Todo) {
	s.r.Add(t)
}

func (s *Srv) Complete(id int) {
	s.r.Complete(id)
}

func (s *Srv) Delete(id int) {
	s.r.Delete(id)
}

func (s *Srv) List() []*entity.Todo {
	return s.r.List()
}

func (s *Srv) SetTodos(todos []*entity.Todo) {
	s.r.todos = todos
}

func (s *Srv) ListNotComplete() []*entity.Todo {
	return s.r.ListNotComplete()
}

func (s *Srv) ListComplete() []*entity.Todo {
	return s.r.ListComplete()
}

func (s *Srv) CountComplete() int {
	return s.r.CountComplete()
}

func (s *Srv) CountNotComplete() int {
	return s.r.CountNotComplete()
}
