package cli

import (
	"time"

	"github.com/ssibrahimbas/cli-todos/src/entity"
	"github.com/ssibrahimbas/cli-todos/src/file_stream"
	"github.com/ssibrahimbas/cli-todos/src/internal"
)

type Cli struct {
	repo *internal.Repo
	srv  *internal.Srv
	fs *file_stream.FileStream
	filename string
	w *Writer
}

type CliConfig struct {
	Repo *internal.Repo
	Srv *internal.Srv
	FileStream *file_stream.FileStream
	FileName string
}

func New(c *CliConfig) *Cli {
	return &Cli{
		repo: c.Repo,
		srv: c.Srv,
		fs: c.FileStream,
		filename: c.FileName,
		w: NewWriter(),
	}
}

func (c *Cli) Run() error {
	err := c.importTasks()
	if err != nil {
		return err
	}
	err = c.checkFlags()
	if err != nil {
		return err
	}
	c.listTasks()
	err = c.exportTasks()
	return err
}

func (c *Cli) checkFlags() error {
	f := NewFlag()
	if f.Add {
		return c.addTask(f.Args[0])
	}
	if f.Complete != -1 {
		return c.completeTask(f)
	}
	if f.Del != -1 {
		return c.deleteTask(f)
	}
	if f.List {
		return c.listTasks()
	}
	return nil
}

func (c *Cli) addTask(t string) error {
	c.srv.Add(&entity.Todo{
		Content: t,
		Done: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	return nil
}

func (c *Cli) completeTask(f *Flag) error {
	c.srv.Complete(f.Complete)
	return nil
}

func (c *Cli) deleteTask(f *Flag) error {
	c.srv.Delete(f.Del)
	return nil
}

func (c *Cli) listTasks() error {
	todos := c.srv.List()
	total := len(todos)
	notComplete := c.srv.CountNotComplete()
	complete := c.srv.CountComplete()
	c.w.WriteTodos(todos, total, notComplete, complete)
	return nil
}

func (c *Cli) importTasks() error {
	d := []*entity.Todo{}
	err := c.fs.Read(c.filename, &d)
	if err != nil {
		return err
	}
	c.srv.SetTodos(d)
	return nil
}

func (c *Cli) exportTasks() error {
	d := c.srv.List()
	return c.fs.Write(c.filename, d)
}