package config

type ConfigModel struct {
	Uri string
	Db  string
}

var Conf = ConfigModel{
	Uri: "mongodb://172.17.0.2:27017",
	Db:  "messages",
}
