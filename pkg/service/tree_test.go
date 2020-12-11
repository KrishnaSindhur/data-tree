package service_test

import (
	"context"
	"testing"

	"github.com/KrishnaSindhur/data-tree/pkg/contract"
	"github.com/KrishnaSindhur/data-tree/pkg/service"

	"github.com/stretchr/testify/assert"
)

func Test_AddData_EmptyShouldBeAbleToGiveEmptyTree(t *testing.T) {
	ctx := context.Background()
	tree := contract.Tree{}
	data := contract.Data{}
	expected := contract.Tree{Level1:contract.Node1{WebReq:0, TimeSpent:0}, Level2:[]contract.Node2{contract.Node2{Country:"", Node1:contract.Node1{WebReq:0, TimeSpent:0}}},
				Level3:[]contract.Node3{contract.Node3{Device:"", Node2:contract.Node2{Country:"", Node1:contract.Node1{WebReq:0, TimeSpent:0}}}}}

	treeData, err := service.AddData(ctx, data, tree)

	assert.NoError(t, err)
	assert.Equal(t, expected, treeData, "expected value is not matching with actual value")
}

func Test_AddData_ShouldBeAbleToRespondData(t *testing.T) {
	ctx := context.Background()
	tree := contract.Tree{}
	dimensions := []contract.Dimensions{}
	metrics := []contract.Metrics{}
	dimensions = append(dimensions, contract.Dimensions{Key: "device", Value: "mobile"})
	dimensions = append(dimensions, contract.Dimensions{Key: "country", Value: "IN"})
	metrics = append(metrics, contract.Metrics{Key: "webreq", Value: 70})
	metrics = append(metrics, contract.Metrics{Key: "timespent", Value: 30})
	data := contract.Data{Dim: dimensions, Met: metrics}

	node1 := contract.Node1{WebReq: 70, TimeSpent: 30}
	var node2 []contract.Node2
	node2 = append(node2, contract.Node2{Country: "IN", Node1:node1})
	var node3 []contract.Node3
	node3 = append(node3, contract.Node3{Device: "mobile", Node2:contract.Node2{Country: "IN", Node1:node1}})
	expectedValue := contract.Tree{Level1: node1, Level2: node2, Level3: node3}

	treeData, err := service.AddData(ctx, data, tree)

	assert.NoError(t, err)
	assert.Equal(t, expectedValue, treeData, "expected value is not matching with actual value")
}