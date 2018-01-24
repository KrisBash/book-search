#!/usr/bin/ruby

require 'rest-client'

#docker run
#cmd = "docker run -d -p 8222:8222 book-api:latest /go/src/app/src/book-api"
#exec = `#{cmd}`


#test query
@url = "http://127.0.0.1:8222/graphql?"
@test_result = true

def test_query(qs, str)
    result = true  
    urlqs = @url  + qs
    res = RestClient.post urlqs,{}
    res_code = res.code
    res_body = res.body
    if res_code == 200
        puts "Request returned: #{res.code}"
    else
        puts "Request failed. Response code #{res.code}"
        result = false
        @test_result = false
    end
    if res_body.include? str
        puts "Response body contains string #{str}"
        puts res_body
    else
        puts "Response body does not contain string #{str}"
        puts res_body
        @test_result = false
    end    
    return result
 end

puts "Testing a single book"
qs = "query={book(isbn:\"9780545010221\") {isbn,authors,page_count,average_rating}}"
result = test_query(qs, "Rowling")

puts "Testing a single book"
qs = "query={book(isbn:\"1491935677\") {isbn,authors,description,image_links}}"
result = test_query(qs,"Kubernetes")

puts "----"
puts "Test result is: #{@test_result}"

#docker stop
#cmd = "cid=`sudo docker ps |grep book-api|head -n 1 |awk '{print $1}'` && sudo docker stop $cid"
#exec = `#{cmd}`

if @test_result == false
    abort("Failures in testing book-api")
end