package status

import (
	"strings"

	"github.com/coreos/etcd/clientv3"
	"github.com/golang/protobuf/proto"
	"github.com/taku-k/polymerase/pkg/base"
	"github.com/taku-k/polymerase/pkg/status/statuspb"
)

func GetNodesInfo(cli *clientv3.Client) *statuspb.NodeInfoMap {
	res, err := cli.KV.Get(cli.Ctx(), base.NodeInfoKey, clientv3.WithPrefix())
	if err != nil {
		return nil
	}
	result := &statuspb.NodeInfoMap{
		Nodes: make(map[string]*statuspb.NodeInfo),
	}
	for _, kv := range res.Kvs {
		n := strings.Split(string(kv.Key), "/")[1]
		info := &statuspb.NodeInfo{}
		if err := proto.Unmarshal(kv.Value, info); err != nil {
			continue
		}
		result.Nodes[n] = info
	}
	return result
}
