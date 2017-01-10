# Storage

![Alt text](anonymous.jpg)


><font color="Black">Supported Type: </font>&nbsp;
 __<font color="Green">S3,Swift,Posix,NFS,Kdfs etc.</font>__

## __<font color="Crimson">1. Build workspace</font>__

```sh
 mkdir -p $GOPATH/src/github.com/flyaways
 cd $GOPATH/src/github.com/flyaways
 git clone  https://github.com/flyaways/storage.git
```

## __<font color="LightSkyBlue">2. Build</font>__

```sh
 cd $GOPATH/src/github.com/flyaways/storage
 export GOBIN=$pwd/bin make
```

## __<font color="LawnGreen">3. Run</font>__

```sh
 cd $GOPATH/src/github.com/flyaways/storage
 ./bin/agent -config=agent.toml
```

## __<font color="Chocolate">4. Test all of agent interface</font>__

><font color="Black">Bucket Method:</font>&nbsp;
 __<font color="Green">PUT,GET, HEAD, DELETE</font>__

><font color="Black">Object Method:</font>&nbsp;
 __<font color="Green">PUT, GET, HEAD, DELETE</font>__

## __<font color="VioletRed">5. References</font>__

* [__Swift__](http://developer.openstack.org/api-ref/object-storage/)
* [__s3__](http://docs.s3.com/docs/master/)

<font color="Black"><h4 align = "center">©2016 flyaways</h4></font>
