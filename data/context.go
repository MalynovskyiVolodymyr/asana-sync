package data

import (
	"asanaSync/common"

	"gopkg.in/mgo.v2"
)

type Context struct {
	DataBaseSession *mgo.Session
}

func (con *Context) Close() {
	con.DataBaseSession.Close()
}

func (con *Context) DBCollection(collectionName string) *mgo.Collection {
	return con.DataBaseSession.DB(common.AppConfig.DataBaseName).C(collectionName)
}

func NewContext() *Context {
	return &Context{
		DataBaseSession: common.GetDBSession().Copy(),
	}
}
