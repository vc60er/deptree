package moduleinfo

// Module is currently a placeholder for upcoming module information, see "go help list"
type Module struct {
	name string
}

func NewModule(name string) Module {
	return Module{name}
}

func (m *Module) Name() string {
	return m.name
}
