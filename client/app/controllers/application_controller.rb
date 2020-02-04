require 'helloworld_service.rb'

class ApplicationController < ActionController::Base
    def say_hello
        helloworld = HelloworldService.new()
        message = helloworld.call_say_hello("sano yo!")
        render json: message.message
    end
end
