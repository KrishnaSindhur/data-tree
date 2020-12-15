package service_test

import (
	"testing"

	"github.com/KrishnaSindhur/data-tree/pkg/contract"
	"github.com/KrishnaSindhur/data-tree/pkg/service"

	"github.com/stretchr/testify/assert"
)

func Test_AddData_AddingEmptyDataForTree(t *testing.T) {
	tree := contract.Node{}
	tree.WebReq = 140
	tree.TimeSpent = 60
	var node2 [] *contract.Node
	node2 = append(node2, &contract.Node{MetaData: "mobile", WebReq: 140, TimeSpent: 60})
	node1 := contract.Node{MetaData: "IN", WebReq: 140, TimeSpent: 60, Children: node2}
	tree.Children = append(tree.Children, &node1)
	data := contract.Data{}

	err := service.AddData(data, &tree)

	assert.NoError(t, err)
}

func Test_AddData_AddingEmptyTree(t *testing.T) {
	var dim []contract.Dimensions
	var met []contract.Metrics
	dim = append(dim, contract.Dimensions{Key: "country", Value: "IN"})
	dim = append(dim, contract.Dimensions{Key: "device", Value: "mobile"})
	met = append(met, contract.Metrics{Key: "webreq", Value: 140})
	met = append(met, contract.Metrics{Key: "timespent", Value: 60})
	data := contract.Data{Dim: dim, Met: met}
	tree := contract.Node{}

	err := service.AddData(data, &tree)

	assert.NoError(t, err)
}

func Test_GetData_GetTreeDataForGivenQuery(t *testing.T) {
	var dim []contract.Dimensions
	dim = append(dim, contract.Dimensions{Key: "country", Value: "IN"})
	query := contract.Query{Data: dim}

	tree := contract.Node{}
	tree.WebReq = 140
	tree.TimeSpent = 60
	var node2 [] *contract.Node
	node2 = append(node2, &contract.Node{MetaData: "mobile", WebReq: 140, TimeSpent: 60})
	node1 := contract.Node{MetaData: "IN", WebReq: 140, TimeSpent: 60, Children: node2}
	tree.Children = append(tree.Children, &node1)

	var dim1 []contract.Dimensions
	var met []contract.Metrics
	dim1 = append(dim1, contract.Dimensions{Key: "country", Value: "IN"})
	met = append(met, contract.Metrics{Key: "webreq", Value: 140})
	met = append(met, contract.Metrics{Key: "timespent", Value: 60})
	expectedData := contract.Data{Dim: dim1, Met: met}

	data, err := service.GetData(query, tree)

	assert.NoError(t, err)
	assert.Equal(t, expectedData,  data, "Unexpected value")
}

func Test_GetData_GetTreeDataForEmptyTree(t *testing.T) {
	var dim []contract.Dimensions
	dim = append(dim, contract.Dimensions{Key: "country", Value: "IN"})
	query := contract.Query{Data: dim}

	data, err := service.GetData(query, contract.Node{})

	assert.NoError(t, err)
	assert.Equal(t, contract.Data{},  data, "Unexpected value")
}