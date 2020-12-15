package service

import (
	"github.com/KrishnaSindhur/data-tree/pkg/contract"
)


func GetData(query contract.Query, tree contract.Node) (contract.Data, error) {
	var country string
	for _, d := range query.Data {
		if d.Key == "country" {
			country = d.Value
		}
	}
	for _, nodeLevel1 := range tree.Children {
		var metrics []contract.Metrics
		if nodeLevel1.MetaData == country {
			metrics = append(metrics, contract.Metrics{Key: "webreq", Value: nodeLevel1.WebReq})
			metrics = append(metrics, contract.Metrics{Key: "timespent", Value: nodeLevel1.TimeSpent})
			var dim []contract.Dimensions
			dim = append(dim, contract.Dimensions{Key: "country", Value: country})
			return contract.Data{Dim: dim, Met: metrics}, nil
		}
	}
	return contract.Data{}, nil
}

func AddData(data contract.Data, tree *contract.Node) error {
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
	if len(tree.Children) == 0 {
		tree.WebReq = webReq
		tree.TimeSpent = timeSpent
		var node2 [] *contract.Node
		node2 = append(node2, &contract.Node{MetaData: device, WebReq: webReq, TimeSpent: timeSpent})
		node1 := contract.Node{MetaData: country, WebReq: webReq, TimeSpent: timeSpent, Children: node2}
		tree.Children = append(tree.Children, &node1)
		return nil
	}

	for _, nodeLevel1 := range tree.Children {
		if nodeLevel1.MetaData == country {
			for _, nodeLevel2 := range nodeLevel1.Children {
				if nodeLevel2.MetaData == device {
					tree.WebReq += webReq
					tree.TimeSpent += timeSpent
					nodeLevel1.WebReq += webReq
					nodeLevel1.TimeSpent += timeSpent
					nodeLevel2.WebReq += webReq
					nodeLevel2.TimeSpent += timeSpent
					return nil
				}
			}
			tree.WebReq += webReq
			tree.TimeSpent += timeSpent
			nodeLevel1.WebReq += webReq
			nodeLevel1.TimeSpent += timeSpent
			var node2 []*contract.Node
			node2 = append(node2, &contract.Node{MetaData: device, WebReq: webReq, TimeSpent: timeSpent})
			nodeLevel1.Children = node2
			return nil
		}
	}
	tree.WebReq += webReq
	tree.TimeSpent += timeSpent
	var node2 [] *contract.Node
	node2 = append(node2, &contract.Node{MetaData: device, WebReq: webReq, TimeSpent: timeSpent})
	node1 := contract.Node{MetaData: country, WebReq: webReq, TimeSpent: timeSpent, Children: node2}
	tree.Children = append(tree.Children, &node1)
	return nil
}