package validator

import (
	"strconv"

	core "example.com/state_transition/core"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func Init(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// two args a string and an integer
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	_, err := strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value")
	}
	return core.Init(stub, args)
}

func Increment(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// two args a string and an integer
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	return core.Increment(stub, args[0])
}

func Decrement(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// two args a string and an integer
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	return core.Decrement(stub, args[0])
}

func Set(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// two args a string and an integer
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	_, err := strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value")
	}
	return core.Set(stub, args)
}

func Query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// one string arg
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	return core.Query(stub, args[0])
}
