nodeModel = "default"

collect "mib2ip" {
  oids "ipInAddrErrors" {}
  oids  "ipFragCreates" {}
  oids  "ipInDiscards" {}
  oids  "ipInReceives" {}
  oids  "ipFragOKs" {}
  oids  "ipInDelivers" {}
  oids  "ipReasmReqds" {}
  oids  "ipOutRequests" {}
  oids  "ipOutNoRoutes" {}
  oids  "ipInHdrErrors" {}
  oids  "ipForwDatagrams" {}
  oids  "ipOutDiscards" {}
  oids  "ipReasmOKs" {}
  oids  "ipInUnknownProtos" {}
}

collect "interface" {
  oids  "ifPhysAddress" {}
  oids  "ifOperStatus"  {}
  oids  "ifPhysAddress"  {}
  oids  "ifPhysAddress" {}
}

collect "sys" {
  oids "sysLocation" {
    title = "System Location"
  }
  oids "sysName" {
    title = "System Name"
  }
  oids "sysDescr" {
    title = "Description"
  }
  oids "sysObjectID"{}
  oids "sysUpTime" {
    title = "Uptimen"
  }
  oids "sysContact" {
    title = "Comtact"
  }
  oids "sysName" {
    title = "System Name"
  }
}

