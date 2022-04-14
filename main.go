package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/marigold-dev/deku-go-interop"
	c "example.com/state_transition/chaincode"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type message struct {
	Action string
	Args   []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func log(message string) {
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m\n", 31, message)
	fmt.Printf(colored)
}

var chaincode c.Chaincode

func init() {
	// note: state machine is not accessible at this point
}

func main() {
	state_transition := func(input []byte) (return_err error) {
		var message message
		log(fmt.Sprintf("State transition received %s", string(input)))
		err := json.Unmarshal(input, &message)
		check(err)

		stub, err := c.NewChaincodeStub(message.Args, &shim.Handler{}, "", "",
			&pb.ChaincodeInput{}, &pb.SignedProposal{})
		if err != nil {
			return errors.New(fmt.Sprintf("error creating newChaincodeStub: %+v", err.Error()))

		}

		switch message.Action {
		case "Init":
			log(fmt.Sprintf("inside Init case"))
			handleShimResponse(chaincode.Init(stub))

		case "Invoke":
			log(fmt.Sprintf("inside Invoke case"))
			handleShimResponse(chaincode.Invoke(stub))
		default:
			return errors.New(fmt.Sprintf("Not supported Action: %v", message.Action))
		}
		return
	}
	deku_interop.Main(state_transition)
}

func handleShimResponse(resp pb.Response) {
	log(fmt.Sprintf("fabric chaincode resp status: %+v", resp.Status))
	log(fmt.Sprintf("fabric chaincode resp message: %+v", string(resp.Message)))
}
