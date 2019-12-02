# Potential Error Scenarios

[Determining Scenarios](https://drive.google.com/file/d/1r5q5i1sATsP510TCInQ6FhPjWbXVMUyt/view?usp=sharing
)

1. Timeout
2. 400 Not Found

Valid Response
3. Invalid Json
4. Valid json response, no error
5. Invalid Error Interface


Example:
curl -v  https://api.mercadolibre.com/countries/BI

```s
HTTP/1.1 404 Not Found
< Content-Type: application/json; charset=utf-8
< Content-Length: 75
< Date: Tue, 19 Nov 2019 14:08:08 GMT
< Cache-Control: max-age=60,stale-while-revalidate=30, stale-if-error=120
< X-Content-Type-Options: nosniff
< X-Request-Id: 2629f8da-6dfc-4abc-870c-e160a23a4bdd
< X-Frame-Options: DENY
< X-XSS-Protection: 1; mode=block
< Access-Control-Allow-Origin: *
< Access-Control-Allow-Headers: Content-Type
< Access-Control-Allow-Methods: PUT, GET, POST, DELETE, OPTIONS
< Access-Control-Max-Age: 86400
< X-Cache: Error from cloudfront
< Via: 1.1 6afc1c7b9e6d4dbe30a0b3eae05d0f9e.cloudfront.net (CloudFront)
< X-Amz-Cf-Pop: EWR52-C4
< X-Amz-Cf-Id: jxr-vmHs68SHVNvIicd7SXXSpBWqRBSVd963o9fryqHd9e61PBJYNA==
< Connection: Keep-Alive
```
