package main

// Uncomment the tests below after implementing the functions in main.go.

// import (
// 	"testing"
//
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
//
// 	pb "github.com/wimspaargaren/go-training-exercises/21.protobuf/2.marshal-unmarshal/pb"
// )
//
// func newTestProduct() *pb.Product {
// 	return &pb.Product{
// 		Id:      1,
// 		Name:    "Gopher Plushie",
// 		Price:   29.99,
// 		InStock: true,
// 	}
// }
//
// func TestMarshalUnmarshal(t *testing.T) {
// 	product := newTestProduct()
//
// 	data, err := MarshalProduct(product)
// 	require.NoError(t, err)
// 	assert.NotEmpty(t, data)
//
// 	result, err := UnmarshalProduct(data)
// 	require.NoError(t, err)
//
// 	assert.Equal(t, product.GetId(), result.GetId())
// 	assert.Equal(t, product.GetName(), result.GetName())
// 	assert.Equal(t, product.GetPrice(), result.GetPrice())
// 	assert.Equal(t, product.GetInStock(), result.GetInStock())
// }
//
// func TestMarshalUnmarshalJSON(t *testing.T) {
// 	product := newTestProduct()
//
// 	data, err := MarshalProductJSON(product)
// 	require.NoError(t, err)
// 	assert.NotEmpty(t, data)
//
// 	result, err := UnmarshalProductJSON(data)
// 	require.NoError(t, err)
//
// 	assert.Equal(t, product.GetId(), result.GetId())
// 	assert.Equal(t, product.GetName(), result.GetName())
// 	assert.Equal(t, product.GetPrice(), result.GetPrice())
// 	assert.Equal(t, product.GetInStock(), result.GetInStock())
// }
