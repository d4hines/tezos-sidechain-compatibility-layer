package core

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func Init(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Printf("key = %s\n", args[0])
	fmt.Printf("value = %s\n", args[1])

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func Increment(stub shim.ChaincodeStubInterface, key string) pb.Response {
	valueBytes, err := stub.GetState(key)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + key + " with error: " + err.Error() + "\"}"
		return shim.Error(jsonResp)
	}

	if valueBytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + key + "\"}"
		return shim.Error(jsonResp)
	}

	var value *int
	err = json.Unmarshal(valueBytes, &value)
	if err != nil {
		return shim.Error("Expecting integer value")
	}

	newValue := *value + 1
	newValueBytes, err := json.Marshal(newValue)

	err = stub.PutState(key, newValueBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func Decrement(stub shim.ChaincodeStubInterface, key string) pb.Response {
	valueBytes, err := stub.GetState(key)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + key + " with error: " + err.Error() + "\"}"
		return shim.Error(jsonResp)
	}

	if valueBytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + key + "\"}"
		return shim.Error(jsonResp)
	}

	var value *int
	err = json.Unmarshal(valueBytes, &value)
	if err != nil {
		return shim.Error("Expecting integer value")
	}

	if *value == 0 {
		return shim.Error("Skipping decrement because current value is 0")
	}

	newValue := *value - 1
	newValueBytes, err := json.Marshal(newValue)

	err = stub.PutState(key, newValueBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func Set(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Printf("key = %s\n", args[0])
	fmt.Printf("value = %s\n", args[1])

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func Query(stub shim.ChaincodeStubInterface, key string) pb.Response {
	// Get the state from the ledger
	valbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + key + " with error: " + err.Error() + "\"}"
		return shim.Error(jsonResp)
	}

	if valbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + key + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + key + "\",\"Amount\":\"" + string(valbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(valbytes)
}
