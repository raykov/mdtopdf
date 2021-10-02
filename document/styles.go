package document

// Styles holds a list of all applied styles at the moment
type Styles struct {
	styles []*Style
}

// Add New style
func (s *Styles) Add(style *Style) {
	s.styles = append(s.styles, style)
}

// Current style
func (s *Styles) Current() *Style {
	return s.styles[len(s.styles)-1]
}

// Pop style from the list
func (s *Styles) Pop() *Style {
	if len(s.styles) < 2 {
		return nil
	}

	last := s.styles[len(s.styles)-1]
	s.styles = s.styles[:len(s.styles)-1]
	return last
}
