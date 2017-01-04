#RUBY SWIFT EXAMPLES

#1 CREATE A CONNECTION
#This creates a connection so that you can interact with the server:

require 'cloudfiles'
username = 'account_name:user_name'
api_key  = 'your_secret_key'

conn = CloudFiles::Connection.new(
        :username => username,
        :api_key  => api_key,
        :auth_url => 'http://objects.dreamhost.com/auth'
)

#2 CREATE A CONTAINER
#This creates a new container called my-new-container

container = conn.create_container('my-new-container')

#3 CREATE AN OBJECT
#This creates a file hello.txt from the file named my_hello.txt

obj = container.create_object('hello.txt')
obj.load_from_filename('./my_hello.txt')
obj.content_type = 'text/plain'

#4 LIST OWNED CONTAINERS
#This gets a list of Containers that you own, and also prints out the container name:

conn.containers.each do |container|
        puts container
end

=begin
The output will look something like this:

mahbuckat1
mahbuckat2
mahbuckat3
=end

#5 LIST A CONTAINER’S CONTENTS
#This gets a list of objects in the container, and prints out each object’s name, the file size, and last modified date:

require 'date'  # not necessary in the next version
container.objects_detail.each do |name, data|
        puts "#{name}\t#{data[:bytes]}\t#{data[:last_modified]}"
end

=begin
The output will look something like this:

myphoto1.jpg 251262  2011-08-08T21:35:48.000Z
myphoto2.jpg 262518  2011-08-08T21:38:01.000Z
=end

#6 RETRIEVE AN OBJECT
#This downloads the object hello.txt and saves it in ./my_hello.txt:

obj = container.object('hello.txt')
obj.save_to_filename('./my_hello.txt')

#7 DELETE AN OBJECT
#This deletes the object goodbye.txt:

container.delete_object('goodbye.txt')

#8 DELETE A CONTAINER¶
#Note The container must be empty! Otherwise the request won’t work!
container.delete_container('my-new-container')