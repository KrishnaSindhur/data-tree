package service

import (
	"context"

	"github.com/KrishnaSindhur/data-tree/pkg/contract"
)

func GetData(ctx context.Context, data contract.Query, tree contract.Tree) (contract.Data, error){
	return contract.Data{}, nil
}

//TODO condense the code
//Todo type safety to move constants or enums
func AddData(ctx context.Context, data contract.Data, tree contract.Tree) (contract.Tree, error) {
	var country, device string
	var webReq, timeSpent int
	for _, d := range data.Dim {
		if d.Key == "country" {
			country = d.Value
		}
		if d.Key == "device" {
			device = d.Value
		}
	}
	for _, m := range data.Met {
		if m.Key == "webreq" {
			webReq = m.Value
		}
		if m.Key == "timespent" {
			timeSpent = m.Value
		}
	}

	for _, l3 := range tree.Level3 {
		if l3.Device == device && l3.Country == country {
			tree = updateData(tree, webReq, timeSpent, l3)
			return tree, nil
		}
	}

	tree = insertData(tree, device, country, webReq, timeSpent)
	return tree, nil
}

//internal function
//to insert data either level2 that is country and then level3 device
// or level 3 that is if given device missing got the country
func insertData(tree contract.Tree, device string, country string, webReq int, timeSpent int)  contract.Tree {
	tree.Level1.WebReq += webReq
	tree.Level1.TimeSpent += timeSpent
	for _, l2 := range tree.Level2 {
		if l2.Country == country {
			l2.WebReq += webReq
			l2.TimeSpent += timeSpent
		}
		tree.Level3 = append(tree.Level3, contract.Node3{Device: device, Node2: l2})
		return tree
	}
	tree.Level2 = append(tree.Level2, contract.Node2{Country: country, Node1: tree.Level1})
	tree.Level3 = append(tree.Level3, contract.Node3{Device: device, Node2: contract.Node2{Country: country, Node1: tree.Level1}})
	return tree
}

// internal function if already present the dimension
func updateData(tree contract.Tree, webReq int, timeSpent int, level3 contract.Node3) contract.Tree{
	tree.Level1.WebReq += webReq
	tree.Level1.TimeSpent += timeSpent
	level3.Node1.TimeSpent += timeSpent
	level3.Node1.WebReq += webReq
	level3.Node2.TimeSpent += timeSpent
	level3.Node2.WebReq += webReq
	return tree
}
