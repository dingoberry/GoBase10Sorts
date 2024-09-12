package patterns

type IPlugin interface {
	Insert()
}

type Plugin struct {
	IPlugin
}

func (Plugin) Insert() {
	println("Insert Plugin")
}

type UsbPlugin struct {
	Plugin
}

func (UsbPlugin) Insert() {
	println("Insert Plugin")
}
