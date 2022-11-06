package cli

import "flag"

type Flag struct {
	Add bool
	Complete int
	Del int
	List bool
	Args []string
}

func NewFlag() *Flag {
	add := flag.Bool("add", false, "add a new task")
	complete := flag.Int("complete", -1, "complete a task")
	del := flag.Int("del", -1, "delete a task")
	list := flag.Bool("list", false, "list all tasks")
	flag.Parse()
	return &Flag{
		Add: *add,
		Complete: *complete,
		Del: *del,
		List: *list,
		Args: flag.Args(),
	}
}