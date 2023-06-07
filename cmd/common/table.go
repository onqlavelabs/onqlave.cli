package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/cli"
)

var (
	mapType = map[string]bool{
		"int": true, "int8": true, "int16": true, "int32": true, "int64": true, "uint": true,
		"uint8": true, "uint16": true, "uint32": true, "uint64": true, "float32": true, "float64": true,
		"string": true, "bool": true, "common.ApplicationId": true, "common.ArxId": true, "*bool": true,
	}
)

type DataTable struct {
	table  table.Model
	Height int
}

// NewDataTable cast data from type any to DataTable struct
func NewDataTable(data any) *DataTable {
	var rows []table.Row
	var rt reflect.Type
	var dataColumns []table.Column
	var jsonColumn []string

	// get data type and check if data is a slice
	isSlice, rt := getDataType(data)
	if isSlice == -1 {
		return nil
	}

	//Set column name base on type of the data
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		str := f.Type.String()
		if mapType[str] {
			dataColumns = append(dataColumns, table.Column{Title: f.Name})
			jsonColumn = append(jsonColumn, strings.TrimSuffix(f.Tag.Get("json"), ",omitempty"))
		}
	}

	// import table rows from slice type data
	if isSlice == 1 {
		row, err := importTableRows(data, jsonColumn)
		if err != nil {
			return nil
		}
		return importDataTable(row, dataColumns, true)
	}

	// declare column for struct type data
	var DataColumns = []table.Column{{Title: "Key"}, {Title: "Value"}}

	// import table rows for struct type data
	rv := reflect.ValueOf(data)
	for _, v := range dataColumns {
		//import field name and field value if value length <= TableViewHeight
		value := fmt.Sprintf("%v", rv.FieldByName(v.Title))
		if len(value) <= TableViewWidth {
			rows = append(rows, table.Row{v.Title, value})
			continue
		}

		//import field name and field value if value length > TableViewHeight
		rows = append(rows, table.Row{v.Title, value[:TableViewWidth]})
		n := len(value) / TableViewWidth
		for i := 1; i < n; i++ {
			rows = append(rows, table.Row{"", value[i*TableViewWidth : (i+1)*TableViewWidth]})
		}
		if len(value)%TableViewWidth != 0 {
			rows = append(rows, table.Row{"", value[n*TableViewWidth:]})
		}
	}

	return importDataTable(rows, DataColumns, false)
}

func (m *DataTable) Init() tea.Cmd { return nil }

func (m *DataTable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m *DataTable) View() string {
	view := lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240"))
	return view.Render(m.table.View()) + "\n"
}

// Render *DataTable into tabular view
func (m *DataTable) Render() {
	//return error if m is nil
	if m == nil {
		RenderCLIOutputError("There was an error rendering data: ", errors.New("render table data failed"))
		return
	}

	if m.Height <= 0 {
		m.Height = TableViewHeight
	}

	// render table data
	m.table.Focus()
	m.table.SetHeight(m.Height)
	m.table.SetStyles(table.Styles{
		Cell:     lipgloss.NewStyle().Padding(0, 1, 0, 1),
		Selected: lipgloss.NewStyle().Bold(true).Foreground(cli.Green),
		Header:   lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).Padding(0, 1, 0, 1).BorderBottom(true),
	})

	_, err := tea.NewProgram(m).Run()
	if err != nil {
		RenderCLIOutputError("There was an error rendering data: ", err)
	}
}

// importTableRows import data and column and return table rows
func importTableRows(data any, columns []string) ([]table.Row, error) {
	var mapData []map[string]interface{}
	var tableRows []table.Row

	//marshal then unmarshal data into map
	dataByte, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(dataByte, &mapData); err != nil {
		return nil, err
	}

	//import rows from data
	for _, v := range mapData {
		var row table.Row
		for _, val := range columns {
			if val == "created_at" || val == "updated_at" {
				parsedTime, err := time.Parse("2006-01-02 15:04:05.999999 -0700 MST", v[val].(string))
				if err == nil && !parsedTime.IsZero() {
					v[val] = parsedTime.Format(time.RFC3339)
				}
			}
			if v[val] == nil {
				row = append(row, "")
				continue
			}
			row = append(row, fmt.Sprintf("%v", v[val]))
		}
		tableRows = append(tableRows, row)
	}

	return tableRows, nil
}

// importDataTable set width, height for table and init DataTable
func importDataTable(rows []table.Row, columns []table.Column, isList bool) *DataTable {
	var width []int
	var tableRows []table.Row
	var tableColumns []table.Column

	// import rows data and calculate width for each column
	for _, row := range rows {
		for i, item := range row {
			if len(width) <= i {
				width = append(width, len(item))
				continue
			}

			if width[i] < len(item) {
				width[i] = len(item)
			}
		}

		tableRows = append(tableRows, row)
	}

	// import columns data and calculate min max six for table columns
	for i, column := range columns {
		columnWidth := TableViewMinColWidth

		if width[i] > columnWidth {
			columnWidth = width[i]
			if width[i] > TableViewMaxColWidth && isList {
				columnWidth = TableViewMaxColWidth
			}
		}

		tableColumns = append(tableColumns, table.Column{Title: column.Title, Width: columnWidth})
	}

	//set data table height
	height := TableViewHeight
	if !isList {
		height = len(tableRows)
	}

	return &DataTable{table: table.New(table.WithColumns(tableColumns), table.WithRows(tableRows)), Height: height}
}

// getDataType return type of input data and input data is slice or not
func getDataType(data any) (int, reflect.Type) {
	// get type of data
	reflectType := reflect.TypeOf(data)

	// data type is slice
	if reflectType.Kind() == reflect.Slice {
		rt := reflectType.Elem()
		if rt.Kind() != reflect.Struct {
			return -1, nil
		}
		return 1, rt
	}

	// data is struct
	if reflectType.Kind() != reflect.Struct {
		return -1, nil
	}
	return 0, reflectType
}
