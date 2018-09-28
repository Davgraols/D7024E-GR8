package d7024e

import "fmt"

var (
	id            = NewRandomKademliaID()
	me            = NewContact(id, "127.0.0.1")
	routingTable  = NewRoutingTable(me)
	bootstrapId   = NewKademliaID("77ff0a3a0ec73e10ff408ece8728f84ae1af7bbf")
	bootstrapNode = NewContact(bootstrapId, "127.0.0.1")
	network       = Network{Port: "4000", BootstrapIP: "127.0.0.1"}
	Requests      = make(chan RPC, 5)
)

func Init() {
	fmt.Println(bootstrapId.String())

	for {
		msg := <-Requests
		fmt.Println(msg.RpcType)
		switch msg.RpcType {
		case 1:
			fmt.Println("PING", msg.SenderIp)
		default:
			fmt.Println("default")
		}
	}

}