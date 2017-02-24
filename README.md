# Storage
## The sotage agent provide object storage api just looklike s3
><font color="Green">Supported Type: </font>&nbsp;
 __<font color="Green">
 	s3
	siwft
	posix
	oss
        qiniu
    </font>__

## __<font color="Black">1. Build workspace</font>__

```sh
 mkdir -p $GOPATH/src/github.com/flyaways
 cd $GOPATH/src/github.com/flyaways
 git clone  https://github.com/flyaways/storage.git
```

## __<font color="Black">2. Build</font>__

```sh
 cd storage
 make
```

## __<font color="Black">3. Run</font>__

```sh
 ./docker/bin/storage -config=storage.toml
```

## __<font color="Black">4. Test all of storage interface</font>__

><font color="Green">Bucket Method:</font>&nbsp;
 __<font color="Green">PUT,GET, HEAD, DELETE,etc</font>__

><font color="Green">Object Method:</font>&nbsp;
 __<font color="Green">PUT, GET, HEAD, DELETE,etc</font>__

><font color="Green">Service Method:</font>&nbsp;
 __<font color="Green">GET</font>__


## __<font color="Black">5. References</font>__

* [__Swift__](http://developer.openstack.org/api-ref/object-storage/)
* [__s3__](http://docs.s3.com/docs/master/)

<font color="Green"><h4 align = "center">©2017 flyaways</h4></font>
