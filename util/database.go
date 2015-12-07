package util

import (
	// "fmt"
	// "log"
	// "os"
	// "github.com/cmu440-F15/paxosapp/paxos"
	"github.com/cmu440-F15/paxosapp/rpc/paxosrpc"
)

type Database struct {
	database map[string]interface{}
	paxos    *Paxos
}

func NewDatabase(paxos *Paxos) *Database {
	db := Database{
		database: make(map[string]interface{}),
		paxos:    paxos,
	}
	return &db
}

func (db *Database) Put(key string, value interface{}) {
	db.database[key] = value
}

func (db *Database) GetValue(key string) (interface{}, bool) {
	val, ok := db.database[key]
	return val, ok
}

func (db *Database) Remove(key string) {
	delete(db.database, key)
}

func (db *Database) PaxosPut(key string, value interface{}) {
	// Get next proposal number
	proposalNumArgs := &paxosrpc.ProposalNumberArgs{Key: key}
	var proposalNumReply paxosrpc.ProposalNumberReply
	if err := db.paxos.conn.Call("PaxosNode.GetNextProposalNumber", proposalNumArgs, &proposalNumReply); err != nil {
		panic(err)
	}

	// Propose value
	proposeArgs := &paxosrpc.ProposeArgs{
		N:   proposalNumReply.N,
		Key: key,
		V:   value,
	}
	var proposeReply paxosrpc.ProposeReply
	if err := db.paxos.conn.Call("PaxosNode.Propose", proposeArgs, &proposeReply); err != nil {
		panic(err)
	}
}

func (db *Database) PaxosGetValue(key string) (interface{}, bool) {
	// Get value using paxos
	getValArgs := &paxosrpc.GetValueArgs{Key: key}
	var getValReply paxosrpc.GetValueReply
	if err := db.paxos.conn.Call("PaxosNode.GetValue", getValArgs, &getValReply); err != nil {
		panic(err)
	}

	if getValReply.Status == paxosrpc.KeyNotFound {
		// fmt.Println("db GetValue err: KeyNotFound")
		return nil, false
	}

	// log.Println(key, getValReply.V)

	return getValReply.V, true
}
