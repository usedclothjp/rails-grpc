require 'helloworld_services_pb.rb'
require 'helloworld_pb.rb'
require 'grpc'

class ApplicationController < ActionController::Base
    def say_hello
        stub = Helloworld::Greeter::Stub.new('localhost:50051', :this_channel_is_insecure)
        stub = Helloworld::Greeter::Stub.new('localhost:50051', :this_channel_is_insecure)
        message = stub.say_hello(Helloworld::HelloRequest.new(name: "tomohiro")).message
        render json: message
    end
end
