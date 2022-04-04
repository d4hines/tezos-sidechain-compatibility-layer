package chaincode

import (
	"fmt"

	validator "example.com/state_transition/validator"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Chaincode struct {
}

func (t *Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("chaincode Init")
	_, args := stub.GetFunctionAndParameters()
	return validator.Init(stub, args)
}

func (t *Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("chaincode Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "Increment" {
		return validator.Increment(stub, args)
	} else if function == "Decrement" {
		return validator.Decrement(stub, args)
	} else if function == "Set" {
		return validator.Set(stub, args)
	} else if function == "Query" {
		return validator.Query(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"Increment\" \"Decrement\" \"Set\" \"Query\"")
}
