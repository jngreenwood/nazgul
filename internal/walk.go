package internal

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gosnmp/gosnmp"
)

func (a *Agent) Walk(node *Node) error {

	if !a.ready.TryLock() {
		return errors.New("agent not ready")
	}

	gosnmp.Default.Target = node.Host
	gosnmp.Default.Port = node.Port
	gosnmp.Default.Community = node.Community
	gosnmp.Default.Version = gosnmp.Version2c
	gosnmp.Default.Timeout = time.Duration(10) * time.Second
	gosnmp.Default.Logger = gosnmp.NewLogger(log.New(os.Stdout, "", 0))

	err := gosnmp.Default.Connect()

	if err != nil {
		return err
	}
	defer gosnmp.Default.Conn.Close()

	err = gosnmp.Default.BulkWalk("1.1", printValue)
	if err != nil {
		return err
	}

	return nil

}

func printValue(pdu gosnmp.SnmpPDU) error {
	fmt.Printf("%s = ", pdu.Name)

	switch pdu.Type {
	case gosnmp.OctetString:
		b := pdu.Value.([]byte)
		fmt.Printf("STRING: %s\n", string(b))
	default:
		fmt.Printf("TYPE %d: %d\n", pdu.Type, gosnmp.ToBigInt(pdu.Value))
	}
	return nil
}
