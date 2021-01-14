# To Test in Mac
```
etcd --enable-v2
etcdctl member list
# 修改    			Endpoints: []string{"http://127.0.0.1:2379"},  指向自己

 etcdctl  put /127.0.0.1 100
 etcdctl get / --prefix --keys-only 


./entropy
nc 127.0.0.1 8888
```