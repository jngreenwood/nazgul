package internal

import (
	"errors"

	"github.com/rs/xid"
)

type Node struct {
	id          string
	Community   string `json:"community"`
	Host        string `json:"host"`
	Active      bool
	Port        uint16 `json:"port"`
	SNMPVersion int
	Model       string `json:"model"`
	Name        string `json:"name"`
}

func NodeFromDefaults(node *Node) *Node {
	return &Node{
		id:          xid.New().String(),
		Community:   "public",
		Active:      true,
		Port:        161,
		SNMPVersion: 2,
		Model:       "default",
	}
}

func (n *Node) GetModel(models map[string]SNMPModel) (*SNMPModel, error) {
	model := models[n.Model]
	if model.NodeModel == "" {
		return nil, errors.New("model not found")
	}
	return &model, nil
}
