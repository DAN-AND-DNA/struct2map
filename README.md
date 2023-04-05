# struct2map
 
## example

```go
package main
type A struct {
	AuthKey     string `protobuf:"bytes,1,opt,name=authKey,proto3" json:"authKey,omitempty"`
	AccountName string `protobuf:"bytes,2,opt,name=accountName,proto3" json:"accountName,omitempty"`
	Age         int    `protobuf:"bytes,3,opt,name=age,proto3" json:"age,omitempty"`
}
/*
convert A to map
	
ret == map[string]interface{} {
            "authKey": "d1234",
            "age":     7,
        }
*/
ret := Struct2Map(&A{AuthKey: "d1234", AccountName: "", Age: 7},  "json", JsonTagNameParser, true)
ret := Struct2Map(A{AuthKey: "d1234", AccountName: "", Age: 7},  "json", JsonTagNameParser, true)
```