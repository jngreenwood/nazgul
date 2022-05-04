package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"sync"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

var (
	ErrNodeNotFound = "Node Node Found"
)

type Agent struct {
	Models map[string]SNMPModel
	Nodes  map[string]Node
	oids   oids
	ready  sync.Mutex
}

func NewAgent() *Agent {
	var agent = &Agent{
		ready:  sync.Mutex{},
		Models: make(map[string]SNMPModel),
		Nodes:  make(map[string]Node),
	}
	//mark that we are not ready yet
	agent.ready.Lock()
	return agent
}

func (a *Agent) Start() {
	//first lets load our models
	var model SNMPModel
	err := hclsimple.DecodeFile("models/default.hcl", nil, &model)
	if err != nil {
		log.Fatalf("Failed to load default model: %s", err)
	}
	a.Models["default"] = model

	//now bootstra the oids
	oids, err := loadOids()
	if err != nil {
		log.Fatalf("Failed to load oids: %s", err)
	}
	a.oids = *oids

	rawNodes, err := ioutil.ReadFile("./nodes.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload []Node
	err = json.Unmarshal(rawNodes, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	for _, rawNode := range payload {
		node := NodeFromDefaults(&rawNode)
		node.Host = rawNode.Host
		node.Community = rawNode.Community
		node.Name = rawNode.Name

		a.Nodes[node.Name] = *node
		log.Printf("Node is %#v", node)
	}

	a.ready.Unlock()
}

func (a *Agent) GetNode(name string) (*Node, error) {
	node := a.Nodes[name]
	if node.id == "" {
		return nil, errors.New(ErrNodeNotFound)
	}
	return &node, nil
}
