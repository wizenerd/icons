package icon

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/wizenerd/color"
)

// Size represents the icon size
type Size uint

//common icon size
const (
	None Size = iota
	Size18
	Size24
	Size36
	Size48
)

func (s Size) String() string {
	o := ""
	switch s {
	case Size18:
		o = "18"
	case Size24:
		o = "24"
	case Size36:
		o = "36"
	case Size48:
		o = "48"
	}
	return o
}

// Icon represents material design icon component
type Icon struct {
	vecty.Core
	Name     Name
	Children vecty.MarkupOrComponentOrHTML
	Size     Size
	Classes  vecty.ClassMap
}

// Render implements vecty.COmponent
func (i *Icon) Render() *vecty.HTML {
	c := make(vecty.ClassMap)
	c["material-icons"] = true
	if i.Size.String() != "" {
		c["md="+i.Size.String()] = true
	}
	if i.Classes != nil {
		for k, v := range i.Classes {
			c[k] = v
		}
	}
	return elem.Italic(
		c,
		i.Children,
		vecty.Text(string(i.Name)),
	)
}

//SetColor sets the icon color to c
func (i *Icon) SetColor(c color.Color) {
	if i.Classes == nil {
		i.Classes = make(vecty.ClassMap)
	}
	i.Classes[c.Text()] = true
}
