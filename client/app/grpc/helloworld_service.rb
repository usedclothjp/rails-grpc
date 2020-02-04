require 'helloworld_services_pb.rb'
require 'helloworld_pb.rb'
require 'grpc'

class HelloworldService
    def initialize()
        @stub = Helloworld::Greeter::Stub.new('localhost:50051', :this_channel_is_insecure)
    end

    def call_say_hello(name)
        @stub.say_hello(Helloworld::HelloRequest.new(name: name))
    end
end
