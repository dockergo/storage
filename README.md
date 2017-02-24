# ** Storage **

## The sotage agent provide object storage api just looklike s3

><font color="Green">Supported: s3 siwft posix oss qiniu</font>

## __<font color="Black">1. Build workspace</font>__

```sh
 $go get github.com/flyaways/storage
```

## __<font color="Black">2. Build</font>__

```sh
 $cd storage
 $make
```

## __<font color="Black">3. Run</font>__

```sh
 $./bin/agent -config=./etc/storage.toml
```

## __<font color="Black">4. Testing</font>__

**<font color=Green>Parameters List</font>**

|Type|PUT|POST|GET|HEAD|DELETE|
|---|---|---|---|---|---|
|Bucket|✔|✔|✔|✔|✔||
|Object|✔|✔|✔|✔|✔||
|Service|✔|✖|✖|✖|✖||

## __<font color="Black">5. References</font>__

* [__Swift__](http://developer.openstack.org/api-ref/object-storage/)
* [__s3__](http://docs.s3.com/docs/master/)

<font color="Green"><h4 align = "center">©2017 flyaways</h4></font>
