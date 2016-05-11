package config

import(
  "hastur/configio"
)

const(
  
)

type AssetConfig struct{
	path []string
}

func (ac *AssetConfig)New(){
  ac.path = []string{
    "/img/",
    "/stylesheet/",
    "/js/",
    "/fonts/",
  }
}

func (ac *AssetConfig)GetValue(key string)interface{}{
  var val interface{}

  switch key{
  case configio.ACPath:
    val = ac.path
  }

  return val
}