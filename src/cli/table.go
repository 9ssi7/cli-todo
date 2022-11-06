package cli

import "github.com/alexeyco/simpletable"

type Table struct {
	t *simpletable.Table
}

func NewTable() *Table {
	return &Table{
		t: simpletable.New(),
	}
}

func (t *Table) SetHeader(header ...string) {
	t.t.Header = &simpletable.Header{}
	cell := []*simpletable.Cell{}
	for _, h := range header {
		cell = append(cell, &simpletable.Cell{
			Align: simpletable.AlignCenter,
			Text:  h,
		})
	}
	t.t.Header.Cells = cell
}

func (t *Table) AddRow(row ...string) {
	cell := []*simpletable.Cell{}
	for _, r := range row {
		cell = append(cell, &simpletable.Cell{
			Text: r,
		})
	}
	t.t.Body.Cells = append(t.t.Body.Cells, cell)
}

func (t *Table) SetFooter(footer ...string) {
	cell := []*simpletable.Cell{}
	for _, f := range footer {
		cell = append(cell, &simpletable.Cell{
			Align: simpletable.AlignCenter,
			Text:  f,
			Span: 5,
		})
	}
	t.t.Footer = &simpletable.Footer{
		Cells: cell,
	}
}

func (t *Table) SetStyle(style *simpletable.Style) {
	t.t.SetStyle(style)
}

func (t *Table) String() string {
	return t.t.String()
}