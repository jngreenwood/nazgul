package internal

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	g "github.com/gosnmp/gosnmp"
)

func (a *Agent) Poll(node *Node) error {

	if !a.ready.TryLock() {
		return errors.New("agent not ready")
	}

	model, err := node.GetModel(a.Models)
	if err != nil {
		return err
	}

	oids, translate := model.translate(&a.oids)

	log.Printf("oids are %#v", oids)

	params := &g.GoSNMP{
		Target:    node.Host,
		Port:      node.Port,
		Community: node.Community,
		Version:   g.Version2c,
		Timeout:   time.Duration(10) * time.Second,
		Logger:    g.NewLogger(log.New(os.Stdout, "", 0)),
	}

	err = params.Connect()

	if err != nil {
		log.Fatalf("Connect() err: %v", err)
		return err
	}
	defer params.Conn.Close()

	result, err := params.Get(oids) // Get() accepts up to g.MAX_OIDS
	if err != nil {
		log.Fatalf("Get() err: %v", err)
		return err
	}

	for _, variable := range result.Variables {
		fmt.Printf("oid: %s name: %s ", variable.Name, translate[variable.Name])

		// the Value of each variable returned by Get() implements
		// interface{}. You could do a type switch...
		switch variable.Type {
		case g.OctetString:
			bytes := variable.Value.([]byte)
			fmt.Printf("string: %s\n", string(bytes))
		default:
			// ... or often you're just interested in numeric values.
			// ToBigInt() will return the Value as a BigInt, for plugging
			// into your calculations.
			fmt.Printf("number: %d\n", g.ToBigInt(variable.Value))
		}
	}
	return nil
}
