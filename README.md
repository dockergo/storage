# **Storage Agent**

### The sotage agent provide object storage api just look like s3.
### Supported: s3 siwft posix oss qiniu

### **<a href="https://flyaways1.gitbooks.io/storage/content/">Storage Agent document</a>**

![agent](agent.png)

# **Local Run**

## __<font color="Black">1. Build workspace</font>__

```sh
 $ go get github.com/flyaways/storage
```

## __<font color="Black">2. Build</font>__

```sh
 $ cd storage
 $ make
```

## __<font color="Black">3. Run</font>__

```sh
 $ ./bin/agent -config=./etc/storage.toml
```

# **Docker Run**

## __<font color="Black">1. Build workspace</font>__

```sh
 $ go get github.com/flyaways/storage
```

## __<font color="Black">2. Build</font>__

```sh
 $ cd storage
 $ ./build.sh
```

## __<font color="Black">3. Run</font>__

```sh
 $ docker run -d --restart=always \
 $       -p 8080:8080 \
 $       -e STORAGE_TYPE=s3 \
 $       -e HTTP_SCHEME=http \
 $       -e S3_ADDR=192.168.20.4 \
 $       -e S3_ADDR_PORT=8888 \
 $       -e S3_ACCESSKEY=ZAZW0PO781UDXLA4HGC7 \
 $       -e S3_SECRETKEY=ORBLBg0P6kdObZ6uudMDEWuiTUPNKMwArNyHWRNu \
 $       agent:latest
```

# **Interface**

### Method Lists

|Type|PUT|POST|GET|HEAD|DELETE|
|---|---|---|---|---|---|
|**Bucket**|✔|✔|✔|✔|✔||
|**Object**|✔|✔|✔|✔|✔||
|**Service**|✔|✖|✖|✖|✖||

## **References**

* [__Swift__](http://developer.openstack.org/api-ref/object-storage/)
* [__s3__](http://docs.s3.com/docs/master/)

<font color="Green"><h4 align = "center">©2017 </h4></font>
