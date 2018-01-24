## Book Search

An example app for fun and experimentation. This has no real-world value as-is.

Organized as a set of microservices for deployment to Kubernetes. Components:

* book-api - a book lookup microservice that returns details of a book by it's ISBN. Incoming requests are implmented in GraphQL and data are retrieved through simple REST calls to Google's book API. 
* book-website - a trivial ReactJS front-end. Books are searched via book-api. 
* cache-db - a REDIS cache service for caching of serialized results.

## Pipeline status:

[![pipeline status](https://gitlab.com/krisbash/book-search/badges/master/pipeline.svg)](https://gitlab.com/krisbash/book-search/commits/master)
