require 'uri'
require 'net/http'

url = URI("https://api.dev.name.com/core/v1/hello")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

request = Net::HTTP::Get.new(url)
request["Authorization"] = 'Basic QWV4YW1wbGUgOk1BbmtpbmQ='

response = http.request(request)
puts response.read_body
