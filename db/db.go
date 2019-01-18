package db

import (
	ns "../names"
)

var AllNames ns.Names

func GetAllNames() (ns.Names) {
	return AllNames
}

func AddName(name *ns.Name) (*ns.Name) {
	AllNames.Name = append(AllNames.Name, name)
	return name
}