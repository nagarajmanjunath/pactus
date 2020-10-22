package stats

import (
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/zarbchain/zarb-go/crypto"
	"github.com/zarbchain/zarb-go/logger"
	"github.com/zarbchain/zarb-go/message"
	"github.com/zarbchain/zarb-go/util"
)

// Stats hold statistic data about peers' behaviors
type Stats struct {
	peers     map[peer.ID]*Peer
	nodes     map[crypto.Address]*Node
	maxHeight int
	logger    *logger.Logger
}

func NewStats(logger *logger.Logger) *Stats {
	return &Stats{
		peers:  make(map[peer.ID]*Peer),
		nodes:  make(map[crypto.Address]*Node),
		logger: logger,
	}
}

func (s *Stats) PeersCount() int {
	return len(s.peers)
}

func (s *Stats) getPeer(peerID peer.ID) *Peer {
	if peer, ok := s.peers[peerID]; ok {
		return peer
	}
	p := NewPeer()
	s.peers[peerID] = p
	return p
}

func (s *Stats) getNode(addr crypto.Address) *Node {
	if node, ok := s.nodes[addr]; ok {
		return node
	}
	n := NewNode()
	s.nodes[addr] = n
	return n
}

func (s *Stats) ParsPeerMessage(peerID peer.ID, msg *message.Message) {
	peer := s.getPeer(peerID)
	node := s.getNode(msg.Initiator)

	peer.receivedMsg = peer.receivedMsg + 1

	//ourHeight, _ := syncer.state.LastBlockInfo()
	switch msg.PayloadType() {
	case message.PayloadTypeStatusReq:
		pld := msg.Payload.(*message.StatusReqPayload)
		s.maxHeight = util.Max(s.maxHeight, pld.Height)

	case message.PayloadTypeBlocksReq:

	case message.PayloadTypeTxRes:
		//pld := msg.Payload.(*message.TxResPayload)

	case message.PayloadTypeTxReq:
		//pld := msg.Payload.(*message.TxReqPayload)

	case message.PayloadTypeHRS:
		pld := msg.Payload.(*message.HRSPayload)
		node.hrs = pld.HRS

	case message.PayloadTypeProposal:
		//pld := msg.Payload.(*message.ProposalPayload)

	case message.PayloadTypeBlock:
		//pld := msg.Payload.(*message.BlockPayload)

	case message.PayloadTypeVote:
		//pld := msg.Payload.(*message.VotePayload)

	case message.PayloadTypeVoteSet:
		//pld := msg.Payload.(*message.VoteSetPayload)

	default:
		s.logger.Error("Unknown message type", "msg", msg)
	}
}