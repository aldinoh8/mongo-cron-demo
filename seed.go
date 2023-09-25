package main

// func main() {
// db := config.InitDb()

// coll := db.Collection("users")
// newUsers := []any{
// 	User{Username: "Robin Van Persie", PhoneNumber: "0828282828"},
// 	User{Username: "Virgil Van Dijk", PhoneNumber: "09219191919"},
// 	User{Username: "Atep", PhoneNumber: "19727219"},
// }
// result, err := coll.InsertMany(context.TODO(), newUsers)
// if err != nil {
// 	panic(err)
// }

// fmt.Println("success create", result.InsertedIDs)

// coll := db.Collection("invoices")
// objId1, _ := primitive.ObjectIDFromHex("6510f6bd18c4c5497ec5a431")
// objId2, _ := primitive.ObjectIDFromHex("6510f6bd18c4c5497ec5a432")
// objId3, _ := primitive.ObjectIDFromHex("6510f6bd18c4c5497ec5a433")

// newInvoice := []any{
// 	Invoice{UserId: objId1, Amount: 50000, Status: "CREATED", ChannelType: "Akulaku"},
// 	Invoice{UserId: objId1, Amount: 25000, Status: "CREATED", ChannelType: "Akulaku"},
// 	Invoice{UserId: objId1, Amount: 30000, Status: "CREATED", ChannelType: "Akulaku"},
// 	Invoice{UserId: objId1, Amount: 33000, Status: "CREATED", ChannelType: "Spaylater"},
// 	Invoice{UserId: objId1, Amount: 15000, Status: "CREATED", ChannelType: "Spaylater"},
// 	Invoice{UserId: objId1, Amount: 65000, Status: "CREATED", ChannelType: "Akulaku"},
// 	Invoice{UserId: objId2, Amount: 50000, Status: "CREATED", ChannelType: "Akulaku"},
// 	Invoice{UserId: objId2, Amount: 25000, Status: "CREATED", ChannelType: "Akulaku"},
// 	Invoice{UserId: objId3, Amount: 15000, Status: "CREATED", ChannelType: "Spaylater"},
// 	Invoice{UserId: objId3, Amount: 65000, Status: "CREATED", ChannelType: "Akulaku"},
// }

// result, err := coll.InsertMany(context.TODO(), newInvoice)
// if err != nil {
// 	panic(err)
// }

// fmt.Println("success create", result.InsertedIDs)
// }
