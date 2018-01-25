#!/usr/bin/ruby

require "net/http"
require "uri"
@test_result = true
@ipAddr = ARGV[0]
if @ipAddr.nil?
    @ipAddr = "127.0.0.1"
end

def test_query(ipAddr)
    result = true  
    url = "http://" + ipAddr
    uri = URI.parse(url)

    puts "Testing request to #{uri}"
    http = Net::HTTP.new(uri.host, uri.port)
    request = Net::HTTP::Get.new(uri.request_uri)
    response = http.request(request)
    puts "Response code is: " + response.code          
    if response.code.to_i > 200
        puts "Unexpected response code!"
        result = false
        @test_result = false
    end
    #response.body   
    return result
 end

test_query(@ipAddr)

 if @test_result == false
    abort("Failures in testing book-website")
end