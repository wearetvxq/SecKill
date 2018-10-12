package setup

import (
	"github.com/wearetvxq/SecKill/sk_layer/config"
	"github.com/wearetvxq/SecKill/sk_layer/logic"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

//初始化Etcd
func InitEtcd(host, productKey string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{host},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Printf("Connect etcd failed. Error : %v", err)
	}

	config.SecLayerCtx.EtcdConf = &config.EtcdConf{
		EtcdConn:          cli,
		EtcdSecProductKey: productKey,
	}

	logic.LoadProductFromEtcd()
}
