package controller

import(
	"net/http"
	"byakhee/http/view"
	"byakhee/http/model"
)

type TopController struct{
}

func (t *TopController) ServeHTTP(w http.ResponseWriter,r *http.Request){
	m := &model.TopModel{}
	m.Init(w,r)
	v := &view.TopView{}
	v.New()
	m.Add(v)
	m.Update()
}
