import unittest
import pytest

from ks3.auth import add_auth_header
from ks3.connection import Connection
from ks3.bucket import Bucket
from ks3.key import Key
from ks3.acl import Policy, ACL, Grant
from ks3.user import User

ak = '1GL02rRYQxK8s7FQh8dV'
sk = '2IDjaPOpFfkq5Zf9K4tKu8k5AKApY8S8eKV1zsRl'
c = Connection(access_key_id=ak, access_key_secret=sk, host='127.0.0.1',port=20808)
bucket_name = 'wpsfiletest'
key_name = '979fe62ccf7db70c6cd62e07cf72c81fee2b5fc4'

'''
#DELETE BUCKET
#c.delete_bucket(bucket_name)
b4 = c.get_bucket(bucket_name)
for k in b4.list():
    k.delete()
c.delete_bucket(bucket_name)


#BUCKET LIST
buckets = c.get_all_buckets()
for b in  buckets:
    print b.name
'''

#PUT BUCKET
b1 = c.create_bucket(bucket_name)

#PUT OBJECT
b2 = c.get_bucket(bucket_name)
k2 = b2.new_key(key_name)
k2.set_contents_from_filename("README.md")

#GET OBJECT
b3 = c.get_bucket(bucket_name)
k3 = b3.get_key(key_name)
k3.get_contents_to_filename("KS3SDK_download_test")