package server

import "html/template"

func Mul(param1 int, param2 int) int {
	return param1 * param2
}

func Add(param1 int, param2 int) int {
	return param1 + param2
}

func Add3(param1, param2, param3 int) int {
	return param1 + param2 + param3
}

func Add4(param1, param2, param3, param4 int) int {
	return param1 + param2 + param3 + param4
}

func Add5(p1, p2, p3, p4, p5 int) int {
	return p1 + p2 + p3 + p4 + p5
}

func Minus(p1, p2 int) int {
	return p1 - p2
}

func Div(p1, p2 uint) float32 {
	return float32(p1) / float32(p2)
}

func Max(p1 uint, pn ...uint) uint {
	max := p1
	for _, p := range pn {
		if p > max {
			max = p
		}
	}
	return max
}

func (s *Server) customTemplateFunctions() {
	s.router.SetFuncMap(template.FuncMap{
		"mul":   Mul,
		"add":   Add,
		"add3":  Add3,
		"add4":  Add4,
		"add5":  Add5,
		"minus": Minus,
		"div":   Div,
		"max":   Max,
	})
}
