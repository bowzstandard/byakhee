package config

import(
	"byakhee/http/controller"
	"hastur/core"
  "hastur/configio"
)

const(

)

type RouterConfig struct{
	path map[string]core.HttpController
}

//個別プロパティの設定必要な時はコンストラクタでInitいれる処理追加で

func(rc *RouterConfig)New(){
  rc.path = map[string]core.HttpController{
    "/":&controller.TopController{},
  }
}

func (rc *RouterConfig)GetValue(key string)interface{}{
  var val interface{}

  switch key{
  case configio.RCPath:
    val = rc.path
  }

  return val
}