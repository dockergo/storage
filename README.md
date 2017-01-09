# Storage Agent

><font color="Black">Supported Type: </font>&nbsp;
 __<font color="Green">S3,Swift,Posix,NFS,Kdfs etc.</font>__

## __<font color="Crimson">1. Build workspace</font>__

```sh
 mkdir -p $GOPATH/src/github.com
 cd $GOPATH/src/github.com
 git clone  https://github.com/flyaways/storage.git
```

## __<font color="Gold">2. Go vendor tool</font>__

```sh
 go get -u -v github.com/kardianos/govendor
 cd $GOPATH/src/github.com/flyaways/storage govendor init
 cd vendor govendor add +external
```

## __<font color="LightSkyBlue">3. Build</font>__

```sh
 cd $GOPATH/src/github.com/flyaways/storage
 export GOBIN=`pwd`/bin make
```

## __<font color="LawnGreen">4. Run</font>__

```sh
 cd $GOPATH/src/github.com/flyaways/storage
 ./bin/agent
```


## __<font color="Chocolate">5. Test all of agent interface</font>__


><font color="Black">Bucket Method:</font>&nbsp;
 __<font color="Green">PUT,GET, HEAD, DELETE</font>__


><font color="Black">Object Method:</font>&nbsp;
 __<font color="Green">PUT, GET, HEAD, DELETE</font>__


## __<font color="VioletRed">6. References</font>__

* [__Swift__](http://developer.openstack.org/api-ref/object-storage/)
* [__Ceph__](http://docs.ceph.com/docs/master/)

<font color="Black"><h4 align = "center">©2016 flyaways</h4></font>
