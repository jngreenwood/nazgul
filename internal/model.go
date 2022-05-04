package internal

import "github.com/hashicorp/hcl/v2/hclsimple"

type SNMPModel struct {
	NodeModel string    `hcl:"nodeModel"`
	Collect   []collect `hcl:"collect,block"`
}

type collect struct {
	Name    string `hcl:"name,label"`
	Indexed bool   `hcl:"indexed,optional"`
	Oids    []oid  `hcl:"oids,block"`
}

type oid struct {
	Oid    string `hcl:"oid,label"`
	Option string `hcl:"option,optional"`
	Title  string `hcl:"title,optional"`
}

func LoadNodeModel(path string) (SNMPModel, error) {
	var model SNMPModel
	err := hclsimple.DecodeFile(path, nil, &model)
	if err != nil {
		return model, err
	}
	return model, nil
}

func (m *SNMPModel) translate(oids *oids) ([]string, map[string]string) {
	var ids = []string{}
	translated := make(map[string]string)
	for _, value := range m.Collect {
		for _, CollectColumns := range value.Oids {
			oid := oids.name2Oid[CollectColumns.Oid]
			if CollectColumns.Option == "" {
				oid += ".0"
			}
			translated[oid] = CollectColumns.Oid
			ids = append(ids, oid)
		}
	}
	return ids, translated
}
