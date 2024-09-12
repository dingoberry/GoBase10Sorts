package patterns

import "fmt"

type content struct {
	Information string
}

type ContentBuilder struct {
	Title   string
	Address string
	Stride  string
}

//goland:noinspection GoExportedFuncWithUnexportedType
func (b ContentBuilder) BuildUp() content {
	content := content{}
	content.Information = fmt.Sprintf("%s\t%s\t%s", b.Title, b.Address, b.Stride)
	return content
}
