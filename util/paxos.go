package util

import (
	"net/rpc"

	"github.com/cmu440-F15/paxosapp/rpc/paxosrpc"
)

type Paxos struct {
	conn     *rpc.Client
	hostport string
}

func NewPaxos(hostport string) *Paxos {
	conn, err := rpc.DialHTTP("tcp", hostport)
	if err != nil {
		panic(err)
	}

	paxos := Paxos{
		conn:     conn,
		hostport: hostport,
	}

	return &paxos
}

// Put value into PaxosNode
func (paxos *Paxos) PaxosPut(key string, value interface{}) {
	// Get next proposal number
	proposalNumArgs := &paxosrpc.ProposalNumberArgs{Key: key}
	var proposalNumReply paxosrpc.ProposalNumberReply
	if err := paxos.conn.Call("PaxosNode.GetNextProposalNumber", proposalNumArgs, &proposalNumReply); err != nil {
		panic(err)
	}

	// Propose value
	proposeArgs := &paxosrpc.ProposeArgs{
		N:   proposalNumReply.N,
		Key: key,
		V:   value,
	}
	var proposeReply paxosrpc.ProposeReply
	if err := paxos.conn.Call("PaxosNode.Propose", proposeArgs, &proposeReply); err != nil {
		panic(err)
	}
}

// GetValue from PaxosNode
func (paxos *Paxos) PaxosGetValue(key string) (interface{}, bool) {
	getValArgs := &paxosrpc.GetValueArgs{Key: key}
	var getValReply paxosrpc.GetValueReply
	if err := paxos.conn.Call("PaxosNode.GetValue", getValArgs, &getValReply); err != nil {
		panic(err)
	}

	if getValReply.Status == paxosrpc.KeyNotFound {
		return nil, false
	}

	return getValReply.V, true
}

// Delete value in PaxosNode
// func (paxos *Paxos) PaxosDelete(key string) {
// 	deleteArgs := &paxosrpc.DeleteArgs{Key: key}
// 	var deleteReply paxosrpc.DeleteReply
// 	if err := paxos.conn.Call("PaxosNode.Delete", deleteArgs, deleteReply); err != nil {
// 		panic(err)
// 	}
// }
