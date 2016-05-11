package config

import(
	"flag"
	"hastur/configio"
)

type ServerConfig struct{
	config map[string]map[string]string
	env string
}

const(
	
)

var env string

func (sc *ServerConfig)New(){
	flag.StringVar(&env,configio.SCFlag,configio.SCEnvDev,"configure on environment")
	flag.Parse()

	sc.config = map[string]map[string]string{
		configio.SCEnvDev:{
			configio.SCSetApp:"localhost:8080",
			configio.SCSetDb:"",
			configio.SCSetAsset:"./public",
			configio.SCSetMail:"localhost:25",
		},
		configio.SCEnvStg:{
			configio.SCSetApp:"localhost:8080",
			configio.SCSetDb:"",
			configio.SCSetAsset:"./public",
			configio.SCSetMail:"",
		},
		configio.SCEnvPro:{
			configio.SCSetApp:"localhost:8080",
			configio.SCSetDb:"",
			configio.SCSetAsset:"./public",
			configio.SCSetMail:"",
		},		
	}

	sc.setEnv(env)
	return
}

func (sc *ServerConfig)setEnv(env string){
	tmp := configio.SCEnvDev
	for key,_:=range sc.config{
		if env==key {
			tmp = env
			break
		}
	}
	sc.env = tmp
	return
}

func (sc *ServerConfig)GetValue(key string) interface{}{
	var val interface{}

	switch key{
	case configio.SCSetApp:
		val = sc.config[sc.env][configio.SCSetApp]
	case configio.SCSetDb:
		val = sc.config[sc.env][configio.SCSetDb]
	case configio.SCSetAsset:
		val = sc.config[sc.env][configio.SCSetAsset]
	case configio.SCFlag:
		val = sc.env
	}

	return val
}

