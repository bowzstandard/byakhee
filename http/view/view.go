package view

import(
	"html/template"
	"net/http"
	"encoding/json"
)

type htmlTemplate struct{
	writer http.ResponseWriter
	path string
	dir string
	data map[string]interface{}
}

func (t *htmlTemplate)init(w http.ResponseWriter,p string,d map[string]interface{}){
	t.writer = w
	t.path = p
	t.dir = "tmpl/"
	t.data = d
}

func (t *htmlTemplate)generate(){
	s,err:=template.ParseFiles(t.dir+t.path)
	if err != nil{
		http.Error(t.writer,err.Error(),http.StatusInternalServerError)
		return
	}
	d := struct{
		Data map[string]interface{}
	}{
		Data:t.data,
	}
	err=s.Execute(t.writer,d)
	if err!=nil{
		http.Error(t.writer,err.Error(),http.StatusInternalServerError)
	}
}

type jsonTemplate struct{
	writer http.ResponseWriter
	data interface{}
}

func (t *jsonTemplate)init(w http.ResponseWriter,d interface{}){
	t.writer = w
	t.data = d
}

func (t *jsonTemplate)generate(){
	j,err:=json.Marshal(t.data)
	if err != nil{
		http.Error(t.writer, err.Error(), http.StatusInternalServerError)
		return
	}
	t.writer.Header().Set("Content-Type", "application/json")
	t.writer.Write(j)
}

