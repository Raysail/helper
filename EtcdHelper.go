package helper

import (
	"errors"
	etcdClient "github.com/coreos/etcd/client"
	"golang.org/x/net/context"
	"time"
)

type EtcdHelper struct {
	etcdKeyApi etcdClient.KeysAPI
}

func NewEtcdHelper(etcdClusterAddress []string) (*EtcdHelper, error) {
	if etcdClusterAddress == nil || len(etcdClusterAddress) == 0 {
		return nil, errors.New("etcd config error")
	}
	cfg := etcdClient.Config{
		Endpoints: etcdClusterAddress,
		Transport: etcdClient.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	client, err := etcdClient.New(cfg)
	if err != nil {
		return nil, err
	}
	etcdHelper := &EtcdHelper{}
	etcdHelper.etcdKeyApi = etcdClient.NewKeysAPI(client)
	return etcdHelper, nil
}

func (etcdHelper *EtcdHelper) GetNodeChildren(key string) ([]EtcdNode, error) {
	resp, err := etcdHelper.etcdKeyApi.Get(context.Background(), key, nil)
	if err != nil {
		return nil, err
	} else {
		if resp.Node.Dir {
			var children []EtcdNode
			for _, single := range resp.Node.Nodes {
				children = append(children, EtcdNode{single.Key, single.Value})
			}
			return children, nil
		} else {
			return nil, errors.New("not a dir")
		}
	}
}
func (etcdHelper *EtcdHelper) RegisterService(key, value string, isUpdate bool) error {
	prevExist := "false"
	if isUpdate {
		prevExist = "true"
	} else {
		prevExist = "false"
	}
	_, err := etcdHelper.etcdKeyApi.Set(context.Background(), key, value, &etcdClient.SetOptions{PrevExist: etcdClient.PrevExistType(prevExist), TTL: 10 * time.Second})
	return err
}

type EtcdNode struct {
	Key   string
	Value string
}
