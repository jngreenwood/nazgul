package internal

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

type oids struct {
	oid2Name map[string]string
	name2Oid map[string]string
}

func (oid *oids) translateOidToName(name string) string {
	return oid.oid2Name[name]
}

func (oid *oids) tranlateNameToOid(name string) string {

	maybe := oid.name2Oid[name]
	if maybe != "" {
		return maybe
	}

	//TODO trim and find again.....
	return ""

}

func loadOids() (*oids, error) {
	oids := oids{
		oid2Name: make(map[string]string),
		name2Oid: make(map[string]string),
	}
	dat, err := os.Open("mibs/nmis_mibs.oid")
	if err != nil {
		return nil, err
	}
	defer dat.Close()
	//fmt.Print(string(dat))
	reader := bufio.NewReader(dat)

	r, err := regexp.Compile(`\"(.*)\".*\"(.*)\"`)
	if err != nil {
		return nil, err
	}
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}
		oidArray := r.FindStringSubmatch(string(line))
		if len(oidArray) == 3 {
			oids.name2Oid[oidArray[1]] = oidArray[2]
			oids.oid2Name[oidArray[2]] = oidArray[1]
		}
	}
	return &oids, nil
}
