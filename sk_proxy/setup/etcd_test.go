package setup

//一会研究下

import (
 	service  "github.com/wearetvxq/SecKill/sk_proxy/config"
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

func TestInitEtcd(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		t.Printf("Connect etcd failed. Error : %v", err)
	}

	var SecInfoConfArr []service.SecProductInfoConf
	SecInfoConfArr = append(
		SecInfoConfArr,
		service.SecProductInfoConf{
			ProductId: 1028,
			StartTime: 0,
			EndTime:   0,
			Status:    0,
			Total:     1000,
			Left:      1000,
		},
		service.SecProductInfoConf{
			ProductId: 1027,
			StartTime: 0,
			EndTime:   0,
			Status:    0,
			Total:     2000,
			Left:      1000,
		},
	)
	data, err := json.Marshal(SecInfoConfArr)
	if err != nil {
		t.Printf("Data Marshal. Error : %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "/oldboy/backend/secskill/product", string(data))
	if err != nil {
		t.Printf("Put failed. Error : %v", err)
	}
	cancel()

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/oldboy/backend/secskill/product")
	if err != nil {
		t.Printf("Get falied. Error : %v", err)
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
