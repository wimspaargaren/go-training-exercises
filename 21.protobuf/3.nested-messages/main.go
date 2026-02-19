package main

// Exercise: Work with nested protobuf messages, repeated fields, and enums.
//
// 1. Run "make generate" to generate the Go code from the proto file.
// 2. Implement the functions below:
//    - CreateAddressBook: build an AddressBook with at least 2 people,
//      each having an address and one or more phone numbers.
//    - CountPhonesByType: count how many phone numbers of a given type
//      exist in the address book.
//    - FindPeopleInCity: return all Person entries whose address matches
//      the given city.
// 3. Validate your solution by running "go test".

// import (
// 	pb "github.com/wimspaargaren/go-training-exercises/21.protobuf/3.nested-messages/pb"
// )

//// CreateAddressBook creates an AddressBook with at least 2 people.
// Each person should have:
// - A name, age, and email
// - At least one phone number with a PhoneType
// - An address with street, city, country, and zip code
// func CreateAddressBook() *pb.AddressBook {
// 	// TODO: implement
// 	return nil
// }

//// CountPhonesByType counts how many phone numbers of the given type
//// exist across all people in the address book.
// func CountPhonesByType(book *pb.AddressBook, phoneType pb.PhoneType) int {
// 	// TODO: implement
// 	return 0
// }

//// FindPeopleInCity returns all people whose address is in the given city.
// func FindPeopleInCity(book *pb.AddressBook, city string) []*pb.Person {
// 	// TODO: implement
// 	return nil
// }

func main() {}
