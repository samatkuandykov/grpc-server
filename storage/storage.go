package storage

import (
	ns "github.com/samatkuandykov/grpc-server/names"
)

var allNames ns.Names

func GetAllNames() (ns.Names) {
	return allNames
}

func AddName(name *ns.Name) (*ns.Name) {
	allNames.Name = append(allNames.Name, name)
	return name
}