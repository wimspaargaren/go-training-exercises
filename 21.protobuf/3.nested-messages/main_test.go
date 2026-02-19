package main

// Uncomment the tests below after implementing the functions in main.go.

// import (
// 	"testing"
//
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
//
// 	pb "github.com/wimspaargaren/go-training-exercises/21.protobuf/3.nested-messages/pb"
// )
//
// func TestCreateAddressBook(t *testing.T) {
// 	book := CreateAddressBook()
// 	require.NotNil(t, book)
// 	assert.GreaterOrEqual(t, len(book.GetPeople()), 2)
//
// 	for _, person := range book.GetPeople() {
// 		assert.NotEmpty(t, person.GetName())
// 		assert.NotEmpty(t, person.GetEmail())
// 		assert.Greater(t, person.GetAge(), int32(0))
// 		assert.NotEmpty(t, person.GetPhones())
// 		assert.NotNil(t, person.GetAddress())
// 		assert.NotEmpty(t, person.GetAddress().GetCity())
// 	}
// }
//
// func TestCountPhonesByType(t *testing.T) {
// 	book := &pb.AddressBook{
// 		People: []*pb.Person{
// 			{
// 				Name: "Alice",
// 				Phones: []*pb.PhoneNumber{
// 					{Number: "123", Type: pb.PhoneType_PHONE_TYPE_MOBILE},
// 					{Number: "456", Type: pb.PhoneType_PHONE_TYPE_WORK},
// 				},
// 			},
// 			{
// 				Name: "Bob",
// 				Phones: []*pb.PhoneNumber{
// 					{Number: "789", Type: pb.PhoneType_PHONE_TYPE_MOBILE},
// 					{Number: "012", Type: pb.PhoneType_PHONE_TYPE_HOME},
// 				},
// 			},
// 		},
// 	}
//
// 	assert.Equal(t, 2, CountPhonesByType(book, pb.PhoneType_PHONE_TYPE_MOBILE))
// 	assert.Equal(t, 1, CountPhonesByType(book, pb.PhoneType_PHONE_TYPE_WORK))
// 	assert.Equal(t, 1, CountPhonesByType(book, pb.PhoneType_PHONE_TYPE_HOME))
// 	assert.Equal(t, 0, CountPhonesByType(book, pb.PhoneType_PHONE_TYPE_UNSPECIFIED))
// }
//
// func TestFindPeopleInCity(t *testing.T) {
// 	book := &pb.AddressBook{
// 		People: []*pb.Person{
// 			{
// 				Name:    "Alice",
// 				Address: &pb.Address{City: "Amsterdam"},
// 			},
// 			{
// 				Name:    "Bob",
// 				Address: &pb.Address{City: "Rotterdam"},
// 			},
// 			{
// 				Name:    "Charlie",
// 				Address: &pb.Address{City: "Amsterdam"},
// 			},
// 		},
// 	}
//
// 	result := FindPeopleInCity(book, "Amsterdam")
// 	assert.Len(t, result, 2)
// 	assert.Equal(t, "Alice", result[0].GetName())
// 	assert.Equal(t, "Charlie", result[1].GetName())
//
// 	result = FindPeopleInCity(book, "Utrecht")
// 	assert.Empty(t, result)
// }
