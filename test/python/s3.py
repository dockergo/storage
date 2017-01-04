#!/bin/python

#PYTHON S3 样例
#1 新建一个连接
'''
下面的代码会新建一个连接，这样你就可以和服务器交互
'''

import boto
import boto.s3.connection
access_key = 'CUPQ9OHNSEQ6KYDLYNK2'
secret_key = 'jmgkfRgE2IkgKH9gmOoboOanID7SNHQ2yNRHw6GF'
server = 'objects.dreamhost.com'

conn = boto.connect_s3(
        aws_access_key_id = access_key,
        aws_secret_access_key = secret_key,
        host = server,
        #is_secure=False,               # uncomment if you are not using ssl
        calling_format = boto.s3.connection.OrdinaryCallingFormat(),
        )

#2 列出用户的所有 BUCKET
'''
下面的代码会列出你的 bucket 的列表。 这也会打印出每个bucket的 bucket 名和创建时间。
'''

for bucket in conn.get_all_buckets():
        print "{name}\t{created}".format(
                name = bucket.name,
                created = bucket.creation_date,
        )

'''
输出形式类似下面这样:
mahbuckat1   2011-04-21T18:05:39.000Z
mahbuckat2   2011-04-21T18:05:48.000Z
mahbuckat3   2011-04-21T18:07:18.000Z
'''

#3 新建一个 BUCKET
'''
下面的代码会新建一个名为 my-new-bucket 的bucket。
'''

bucket = conn.create_bucket('my-new-bucket')

#4 列出 BUCKET 的内容
'''
下面的代码会输出 bucket 内的所有对象列表。 这也会打印出每一个对象的名字、文件尺寸和最近修改时间。
'''

for key in bucket.list():
        print "{name}\t{size}\t{modified}".format(
                name = key.name,
                size = key.size,
                modified = key.last_modified,
                )

'''
输出形式类似下面这样:
myphoto1.jpg 251262  2011-08-08T21:35:48.000Z
myphoto2.jpg 262518  2011-08-08T21:38:01.000Z
'''

#5 删除 BUCKET
'''
Note Bucket必须为空！否则它不会工作!
'''

conn.delete_bucket(bucket.name)


#6 新建一个对象
'''
下面的代码会新建一个内容是字符串``”Hello World!”`` 的文件 hello.txt。
'''

key = bucket.new_key('hello.txt')
key.set_contents_from_string('Hello World!')


#7 修改一个对象的 ACL
'''
下面的代码会将对象 hello.txt 的权限变为公开可读，而将 secret_plans.txt 的权限设为私有。
'''

hello_key = bucket.get_key('hello.txt')
hello_key.set_canned_acl('public-read')
plans_key = bucket.get_key('secret_plans.txt')
plans_key.set_canned_acl('private')

#8 下载一个对象 (到文件)
'''
下面的代码会下载对象 perl_poetry.pdf 并将它存到位置 C:\Users\larry\Documents
'''

key = bucket.get_key('perl_poetry.pdf')
key.get_contents_to_filename('/home/larry/documents/perl_poetry.pdf')

#9 删除一个对象
'''
下面的代码会删除对象 goodbye.txt
'''

bucket.delete_key('goodbye.txt')

#10 生成对象的下载 URLS (带签名和不带签名)
'''
下面的代码会为 hello.txt 生成一个无签名为下载URL。 这个操作是生效是因为前面我们已经设置 hello.txt 的 ACL 为公开可读。
下面的代码同时会为 secret_plans.txt 生成一个有效时间是一个小时的带签名的下载 URL。
带签名的下载 URL 在这个时间内是可用的，即使对象的权限是私有(当时间到期后 URL 将不可用)。
'''

hello_key = bucket.get_key('hello.txt')
hello_url = hello_key.generate_url(0, query_auth=False, force_http=True)
print hello_url

plans_key = bucket.get_key('secret_plans.txt')
plans_url = plans_key.generate_url(3600, query_auth=True, force_http=True)
print plans_url

'''
输出形式类似下面这样:
http://objects.dreamhost.com/my-bucket-name/hello.txt
http://objects.dreamhost.com/my-bucket-name/secret_plans.txt? \
Signature=XXXXXXXXXXXXXXXXXXXXXXXXXXX&Expires=1316027075&AWSAccessKeyId=XXXXXXXXXXXXXXXXXXX
'''