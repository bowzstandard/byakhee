package model

import(
	"hastur/core"
	//"net/http"
)

func performRender(m core.HttpModel,o map[core.HttpView]int){
	for key,_:=range o{
		key.Render(m)
	}
}

func add(o map[core.HttpView]int,v core.HttpView){
	tmp := o
	tmp[v]=1
}

func checkOverlap(o map[core.HttpView]int,v core.HttpView) bool{
	for key,_:=range o{
		if key==v{
			return false
			break
		}
	}
	return true
}

func remove(o map[core.HttpView]int,v core.HttpView){
	tmp := o
	_,ok := tmp[v]
	if ok == false{
		return
	}
	delete(tmp,v)
}

func validateUser() bool{
	return true	
}


