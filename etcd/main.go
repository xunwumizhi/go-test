package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

func main() {
	electionTest()

	leaseTxn()
	leaseTest()
	watchTest()
	kvInOp()
	kvOp()
}

const prefix = "/election-demo"
const prop = "local"

type worker struct {
	name       string
	leaderFlag bool
}

func (w *worker) doCrontab() {
	if w.leaderFlag == true {
		fmt.Println(w.name, "doCrontab")
	}
}

func electionTest() {
	endpoints := []string{"127.0.0.1:2379"}
	donec := make(chan struct{})
	go func() {
		<-time.After(20 * time.Second)
		close(donec)
	}()

	stop := make(chan struct{})
	go func() {
		<-time.After(10 * time.Second)
		close(stop)
	}()

	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	bs, _ := json.Marshal(cli)
	if cli == nil {
		fmt.Println("create clinetv3 error")
	}
	fmt.Println("clientv3: ", string(bs))
	defer cli.Close()

	w1 := &worker{name: "w1"}
	w2 := &worker{name: "w2"}
	go w1.campaign(cli, prefix, prop, donec)
	go w2.campaign(cli, prefix, prop, stop)

	go func() {
		ticker := time.NewTicker(time.Duration(5) * time.Second)
		for {
			select {
			case <-ticker.C:
				go w1.doCrontab()
				go w2.doCrontab()
			}
		}
	}()

	time.Sleep(25 * time.Second)
	fmt.Println("main exit")
}

func (w *worker) campaign(c *clientv3.Client, keyPrifex string, keyValue string, stop <-chan struct{}) {
	// session 回话时长15s，session会自动为lease续约
	fmt.Println("begin newsession")
	s, err := concurrency.NewSession(c, concurrency.WithTTL(15))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("begin newelection")
	e := concurrency.NewElection(s, keyPrifex)

	ctx, ctxCancel := context.WithCancel(context.TODO())
	defer ctxCancel()

	fmt.Println("start campaigning")
	// 未当选leader前，会一直阻塞在Campaign调用
	if err = e.Campaign(ctx, keyValue); err != nil {
		fmt.Println(err)
		return
	}
	log.Println(w.name, "elect: success")
	w.leaderFlag = true

	select {
	case <-stop:
	case <-s.Done(): // 等待lease过期，但由于session为lease自动续约，过期时意味着ctx关闭，session退出，到期后没续约
		w.leaderFlag = false
		log.Println(w.name, "elect: expired")
	}

	log.Println(w.name, "resign...")
	e.Resign(context.TODO()) // 传递一个新ctx
	w.leaderFlag = false

}

func kvOp() {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		kv      clientv3.KV
		getResp *clientv3.GetResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	// 建立一个客户端
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	// 用于读写etcd的键值对
	kv = clientv3.NewKV(client)

	// 写入
	kv.Put(context.TODO(), "name1", "lesroad")
	kv.Put(context.TODO(), "name2", "haha")

	// 读取name为前缀的所有key
	if getResp, err = kv.Get(context.TODO(), "name", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getResp.Kvs)

	// 删除name为前缀的所有key
	if _, err = kv.Delete(context.TODO(), "name", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
		return
	}

	if getResp, err = kv.Get(context.TODO(), "name", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getResp.Kvs)
}

func kvInOp() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
		kv     clientv3.KV
		putOp  clientv3.Op
		getOp  clientv3.Op
		opResp clientv3.OpResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	// 建立一个客户端
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}
	kv = clientv3.NewKV(client)

	// 创建Op :operator
	putOp = clientv3.OpPut("myK", "myV")
	// 执行Op 用kv.Do取代 kv.Put kv.Get...
	if opResp, err = kv.Do(context.TODO(), putOp); err != nil {
		fmt.Println(err)
		return
	}
	bs, _ := json.Marshal(opResp.Put())
	fmt.Println("put opResp: ", string(bs))

	// 创建Op
	getOp = clientv3.OpGet("myK")
	// 执行Op
	if opResp, err = kv.Do(context.TODO(), getOp); err != nil {
		fmt.Println(err)
		return
	}
	bs, _ = json.Marshal(opResp.Get())
	fmt.Println("get opResp: ", string(bs))

}

func leaseTest() {
	var (
		config         clientv3.Config
		client         *clientv3.Client
		err            error
		lease          clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId        clientv3.LeaseID
		putResp        *clientv3.PutResponse
		kv             clientv3.KV
		getResp        *clientv3.GetResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	// 建立一个客户端
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	// 申请一个lease(租约)
	lease = clientv3.NewLease(client)
	// 申请一个10秒的lease
	if leaseGrantResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}
	leaseId = leaseGrantResp.ID

	kv = clientv3.NewKV(client)
	// put一个kv，让它与租约关联起来
	if putResp, err = kv.Put(context.TODO(), "name", "lbwnb", clientv3.WithLease(leaseId)); err != nil {
		fmt.Println(err)
		return
	}
	bs, _ := json.Marshal(putResp)
	fmt.Println("putResp: ", string(bs))

	// 定时看下key过期了没有
	for {
		if getResp, err = kv.Get(context.TODO(), "name"); err != nil {
			fmt.Println(err)
			return
		}
		bs, _ := json.Marshal(getResp)
		fmt.Println("getResp: ", string(bs))

		if getResp.Count == 0 {
			fmt.Println("kv过期")
			break
		}
		time.Sleep(2 * time.Second)
	}
}

func watchTest() {
	var (
		config             clientv3.Config
		client             *clientv3.Client
		err                error
		kv                 clientv3.KV
		watchStartRevision int64
		watcher            clientv3.Watcher
		watchRespChan      <-chan clientv3.WatchResponse
		watchResp          clientv3.WatchResponse
		event              *clientv3.Event
	)

	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	// 建立一个客户端
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}
	// 用于读写etcd的键值对
	kv = clientv3.NewKV(client)

	// 模拟etcd中kv的变化
	go func() {
		for {
			kv.Put(context.TODO(), "name", "lesroad")

			kv.Delete(context.TODO(), "name")

			time.Sleep(1 * time.Second)
		}
	}()

	// 创建一个监听器
	watcher = clientv3.NewWatcher(client)

	// 启动监听 5秒后关闭
	ctx, cancelFunc := context.WithCancel(context.TODO())
	time.AfterFunc(5*time.Second, func() {
		cancelFunc()
	})
	watchRespChan = watcher.Watch(ctx, "name", clientv3.WithRev(watchStartRevision))

	// 处理kv变化事件
	for watchResp = range watchRespChan {
		for _, event = range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("事件put：", string(event.Kv.Value))
			case mvccpb.DELETE:
				fmt.Println("事件delete", string(event.Kv.Key))
			}
		}
	}
}

func leaseTxn() {
	var (
		config         clientv3.Config
		client         *clientv3.Client
		err            error
		kv             clientv3.KV
		lease          clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId        clientv3.LeaseID
		keepRespChan   <-chan *clientv3.LeaseKeepAliveResponse
		keepResp       *clientv3.LeaseKeepAliveResponse
		ctx            context.Context
		cancelFunc     context.CancelFunc
		txn            clientv3.Txn
		txnResp        *clientv3.TxnResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	// 建立一个客户端
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	// lease实现锁自动过期
	// op操作
	// txn事务: if else then

	// 1 上锁 创建租约 自动续租 拿着租约去抢占一个key
	lease = clientv3.NewLease(client)
	// 申请一个5秒的lease
	if leaseGrantResp, err = lease.Grant(context.TODO(), 5); err != nil {
		fmt.Println(err)
		return
	}
	// 拿到租约id
	leaseId = leaseGrantResp.ID
	defer lease.Revoke(context.TODO(), leaseId) // 释放租约

	// 准备一个用于取消自动续租的context
	ctx, cancelFunc = context.WithCancel(context.TODO())
	defer cancelFunc() // 取消续租
	if keepRespChan, err = lease.KeepAlive(ctx, leaseId); err != nil {
		fmt.Println(err)
		return
	}
	// 处理续租应答的协程
	go func() {
		select {
		case keepResp = <-keepRespChan:
			if keepRespChan == nil {
				fmt.Println("租约失效")
				goto END
			} else {
				fmt.Println("收到自动续租应答", keepResp)
			}
		}
	END:
	}()

	txnHandle := func(name string) {
		kv = clientv3.NewKV(client)
		txn = kv.Txn(context.TODO())
		// 定义事务
		// 如果key不存在
		txn.If(clientv3.Compare(clientv3.CreateRevision("mutex"), "=", 0)).
			Then(clientv3.OpPut("mutex", "yes", clientv3.WithLease(leaseId))).
			Else(clientv3.OpGet("mutex"))

		// 提交事务
		if txnResp, err = txn.Commit(); err != nil {
			fmt.Println(name, ": commit 有误", err)
			return
		}

		// 判断是否提交成功
		if !txnResp.Succeeded {
			bs, _ := json.Marshal(txnResp)
			fmt.Println(name, ": commit成功，返回结果：失败", string(bs))
			return
		}
		fmt.Println(name, "提交响应结果：成功")
	}

	go txnHandle("g1")
	go txnHandle("g2")
	time.Sleep(3 * time.Second)
	cancelFunc()
	fmt.Println("取消自动续约，等待租约到期")
	time.Sleep(6 * time.Second)
	txnHandle("g3")

}
