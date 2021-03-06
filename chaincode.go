package greeting

import (
	"fmt"
	"github.com/s7techlab/cckit/extensions/owner"
	"github.com/s7techlab/cckit/router"
	"github.com/s7techlab/cckit/state"
	"github.com/s7techlab/cckit/state/mapping"
)

func NewCC() (*router.Chaincode, error) {
	r := router.New("greeting")
	r.Init(owner.InvokeSetFromCreator)

	cc := NewHelloChaincode()
	if err := RegisterHelloServiceChaincode(r, cc); err != nil {
		return nil, err
	}

	return router.NewChaincode(r), nil
}

func NewHelloChaincode() *HelloServiceChaincodeImpl {
	return &HelloServiceChaincodeImpl{}
}

type HelloServiceChaincodeImpl struct {
}

func (cc *HelloServiceChaincodeImpl) SayHello(ctx router.Context, req *HelloRequest) (*HelloResponse, error) {
	res := HelloResponse{
		Reply: fmt.Sprintf("%v, world!", req.Greeting),
	}

	return &res, nil
}

func (cc *HelloServiceChaincodeImpl) state(ctx router.Context) mapping.MappedState {
	return mapping.WrapState(ctx.State(), mapping.StateMappings{})
}

func (cc *HelloServiceChaincodeImpl) event(ctx router.Context) state.Event {
	return mapping.WrapEvent(ctx.Event(), mapping.EventMappings{}.
		Add(&HelloRequest{}))
}
