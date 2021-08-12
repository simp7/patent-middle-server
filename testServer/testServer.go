package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/logger"
	"io"
)

//go:embed xmlData/*
var fs embed.FS

type testServer struct {
	*gin.Engine
	fileDict map[string]string
}

func NewTestServer() *testServer {

	t := &testServer{Engine: gin.Default()}

	t.fileDict = make(map[string]string)
	t.fileDict["3진법*반도체*누설전류*컴퓨터*!생물학"] = "data1.xml"

	return t
}

func (t *testServer) Start() error {

	t.GET("/", t.Search)

	return t.Run(":8080")

}

func (t *testServer) Search(ctx *gin.Context) {

	fileName := t.fileDict[ctx.Query("word")]

	result, err := fs.Open("xmlData/" + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(ctx.Writer, result)

}

func main() {
	t := NewTestServer()
	if err := t.Start(); err != nil {
		logger.Error(err)
	}
}
