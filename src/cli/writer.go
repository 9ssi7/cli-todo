package cli

import (
	"fmt"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/ssibrahimbas/cli-todos/src/entity"
)

type Writer struct {
	c *ColorImpl
}

func NewWriter() *Writer {
	return &Writer{
		c: NewColor(),
	}
}

func (w *Writer) WriteTodos(d []*entity.Todo, total int, complete int, notComplete int) {
	t := NewTable()
	t.SetHeader("ID", "Content", "Done", "Created At", "Updated At")
	for i, todo := range d {
		tsk := w.c.Blue(todo.Content)
		done := w.c.Blue("✗")
		if todo.Done {
			tsk = w.c.Green(todo.Content)
			done = w.c.Green("✔")
		}
		t.AddRow(fmt.Sprint(i), tsk, done, todo.CreatedAt.Format(time.RFC822), todo.UpdatedAt.Format(time.RFC822))
	}
	ttl := w.c.Gray(fmt.Sprintf("Total: %d", total))	
	nc := w.c.Red(fmt.Sprintf("Not Complete: %d", notComplete))
	c := w.c.Blue(fmt.Sprintf("Complete: %d", complete))
	txt := fmt.Sprintf("%s, %s, %s", ttl, nc, c)
	t.SetFooter(txt)
	t.SetStyle(simpletable.StyleDefault)
	println(t.String())
}