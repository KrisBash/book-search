#!/usr/bin/ruby

require 'rest-client'
require 'open-uri'

#test query

@ipAddr = ARGV[0]
if @ipAddr.nil?
    @ipAddr = "127.0.0.1"
end

def test_query(ipAddr)
    result = true  
    uri = "http://" + ipAddr
    puts "Testing request to #{uri}"
    content = open(uri).read
    puts content
    return result
 end

 test_query(@ipAddr)

#if @test_result == false
 #   abort("Failures in testing book-api")
#end