package model

var ViewFunc func(Model) string

func (m Model) View() string {
	if ViewFunc != nil {
		return ViewFunc(m)
	}
	return ""
}