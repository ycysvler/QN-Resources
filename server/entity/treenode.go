package entity

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type TreeNode struct {
	Id         string `id`
	Name       string `file name`
	Path       string `file path`
	Type       int    `0:dir 1:file`
	CreateTime time.Time
	Children   []TreeNode
}

var root TreeNode

var FILE_PATH = "./data"

func init() {
	_, err := os.Stat(FILE_PATH)
	if err == nil {
		load()

	} else {
		root = TreeNode{Id: "root", Name: "root", Path: "root", Type: 0, CreateTime: time.Now(), Children: []TreeNode{}}
		save()
	}
}

func Instance() *TreeNode {
	return &root
}

func (this *TreeNode) FindNodesById(Id string) []TreeNode {
	if this.Id == Id {
		return this.Children
	} else {
		for _, node := range this.Children {
			return node.FindNodesById(Id)
		}

		return make([]TreeNode, 0)
	}
}

func (this *TreeNode) AddTreeNode(Name string, Path string, Type int) {
	node := TreeNode{}
	node.Id = strconv.FormatInt(time.Now().Unix(), 10)
	node.Name = Name
	node.Path = Path
	node.Type = Type
	node.CreateTime = time.Now()
	this.Children = append(this.Children, node)
	save()
}

func save() {
	var buf bytes.Buffer           // 构建一个buf
	encode := gob.NewEncoder(&buf) // 构建一个encodeer
	encode.Encode(root)            // 序列化数据

	file, err := os.OpenFile(FILE_PATH, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()      // 关闭文件指针
	file.Write(buf.Bytes()) // 写入文件
}

func load() {
	buf, err := ioutil.ReadFile(FILE_PATH)
	if err != nil {
		fmt.Println(err)
	}

	decoder := gob.NewDecoder(bytes.NewReader(buf))
	decoder.Decode(&root)
}
