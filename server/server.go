package server

import(
	"net/http"
	"regexp"
	"hastur/configio"
	"byakhee/config"
	"hastur/core"
	//"byakhee/logio"
	"hastur/session"
)

var env string

const(
	COOKIE_NAME="b_log"
	COOKIE_LIFETIME=3600
)

func Exec(){
	
	//config系
	org:=configio.Call()

	sc:=&config.ServerConfig{}
	sc.New()

	ac:=&config.AssetConfig{}
	ac.New()

	rc:=&config.RouterConfig{}
	rc.New()

	org.RegisterConfig(configio.HCOServer,sc)
	org.RegisterConfig(configio.HCOAsset,ac)
	org.RegisterConfig(configio.HCORoute,rc)

	//session管理
	session.Init(COOKIE_NAME,COOKIE_LIFETIME)
	
	setupRouter()
	setupAssetFS()

	app:=sc.GetValue(configio.SCSetApp).(string)
	http.ListenAndServe(app,nil)
}

func setupRouter(){

	org:=configio.Call()
	rc,err:=org.GetConfig(configio.HCORoute)
	if err!=nil{
		return
	}

	for key,val:= range rc.GetValue(configio.RCPath).(map[string]core.HttpController){
		http.HandleFunc(key,makeHttp(key,val))
	}

}

func setupAssetFS(){
	org:=configio.Call()
	ac,err:=org.GetConfig(configio.HCOAsset)
	if err!=nil{
		return
	}
	sc,err2:=org.GetConfig(configio.HCOServer)
	if err2!=nil{
		return
	}

	dir:=sc.GetValue(configio.SCSetAsset).(string)
	for _,val:= range ac.GetValue(configio.ACPath).([]string){
		http.Handle(val, http.FileServer(http.Dir(dir)))
	}
}

func makeHttp(route string,ctrl core.HttpController) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		path:=regexp.MustCompile("^"+route+"$")
		m:=path.FindStringSubmatch(r.URL.Path)
		if m==nil{
			http.NotFound(w,r)
			return
		}
		ctrl.ServeHTTP(w,r)
	}
}

