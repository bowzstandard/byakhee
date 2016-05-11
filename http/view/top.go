package view

import(
	"hastur/core"
)

type TopView struct{}

func (t *TopView) New(){
}

func (t *TopView) Render(s core.HttpModel){

	p:=&htmlTemplate{}
	p.init(s.GetWriter(),
		"top/index.html",
		nil)
	p.generate()
}


