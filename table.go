package fyne_extend

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewTable(items []TableItem, rowCount func() int) *widget.Table {
	table := widget.NewTableWithHeaders(func() (rows int, cols int) {
		return rowCount(), len(items)
	}, func() fyne.CanvasObject {
		return container.NewHBox(
			widget.NewLabel("Will be replaced"),
			widget.NewHyperlink("", nil),
		)
	}, func(id widget.TableCellID, object fyne.CanvasObject) {
		value := items[id.Col].Data(id.Row, id.Col)
		switch items[id.Col].Type {
		case TableItemHyperLink:
			object.(*fyne.Container).Objects[0].(*widget.Label).SetText(value)
			object.(*fyne.Container).Objects[0].(*widget.Label).Hide()

			targetURL, _ := url.Parse(value)

			object.(*fyne.Container).Objects[1].(*widget.Hyperlink).SetText(value)

			object.(*fyne.Container).Objects[1].(*widget.Hyperlink).SetURL(targetURL)
		case TableItemText:
			object.(*fyne.Container).Objects[0].(*widget.Label).SetText(value)
			object.(*fyne.Container).Objects[1].(*widget.Hyperlink).Hide()
		default:

		}
	})

	table.UpdateHeader = func(id widget.TableCellID, template fyne.CanvasObject) {
		if id.Row != -1 {
			template.(*widget.Label).SetText(fmt.Sprintf(`%d`, id.Row+1))
			return
		}

		if id.Col == 0 {
			template.(*widget.Label).SetText(items[0].Title)
		} else {
			template.(*widget.Label).SetText(items[1].Title)
		}
	}

	for i, item := range items {
		if item.Width > 0 {
			table.SetColumnWidth(i, item.Width)
		}
	}

	return table
}

type TableItem struct {
	Title string
	Type  TableItemType
	Width float32
	Data  func(row, column int) string
}

type TableItemType int

const (
	TableItemText TableItemType = iota + 1
	TableItemHyperLink
)
