package service_test

import (
	"context"
	"testing"

	"github.com/KrishnaSindhur/data-tree/pkg/contract"
	"github.com/KrishnaSindhur/data-tree/pkg/service"
	"github.com/stretchr/testify/assert"
)

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

	err := service.AddData(ctx, data, tree)
	assert.NoError(t, err)
}