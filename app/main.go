package main

import (
  //"os"
  "github.com/gocql/gocql"
)


func main() {
  // connect to the cluster
  cluster := gocql.NewCluster("cassandra")
  cluster.Keyspace = "neurose1"
  cluster.ProtoVersion = 3
  cluster.Consistency = gocql.Quorum
  session,_ := cluster.CreateSession()
  defer session.Close()

  // Vide rest.go
  InitRESTMap(session)
}
