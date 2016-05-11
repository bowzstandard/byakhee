package model

import(
	"hastur/core"
	"hastur/logio"
	"net/http"
	"strings"
	"hastur/session"
)

type TopModel struct{
	writer http.ResponseWriter
	request *http.Request
	manager core.SessionManager
	observers map[core.HttpView]int
	data []core.DataModel
	hasSession bool
	error core.LogIO
}

func (s *TopModel)Init(w http.ResponseWriter,r *http.Request){
	s.writer = w
	s.request = r
	s.observers = make(map[core.HttpView]int)
	s.hasSession = false
	s.manager = session.Call()
}

func (s *TopModel) PerformRender(){
	performRender(s,s.observers)
}

func (s *TopModel) Add(o core.HttpView){
	if !checkOverlap(s.observers,o){
		return
	}
	add(s.observers,o)
}

func (s *TopModel) Remove(o core.HttpView){
	remove(s.observers,o)
}

func (s *TopModel) Update(){

	defer s.PerformRender()

	pl,err:=s.manager.SessionStart(s.writer,s.request)

	if err!=nil{
		l:=&logio.HttpErrorLogIO{}
		l.New(
			logio.ServerLog,
			logio.SessionFailed,
			s.request.URL.Path,
			err.Error()+"invalid session!-> request ip:"+strings.Split(s.request.RemoteAddr,":")[0],
		)
		s.error = l
		return
	}
	
	if pl!=nil{
		s.hasSession = true
	}

	return
	
}

func (s *TopModel) GetObservers() map[core.HttpView]int{
	return s.observers
}

func (s *TopModel) GetWriter() http.ResponseWriter{
	return s.writer
}

func (s *TopModel) GetRequest() *http.Request{
	return s.request
}

func (s *TopModel) HasError()core.LogIO{
	return s.error
}

func (s *TopModel) HasSession() bool{
	return s.hasSession
}

func (s *TopModel) GetValue(key string) interface{}{
	return ""
}

