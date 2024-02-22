Steps to test:
cd in the root directory
run "docker-compose up -d" command

Send a POST Request to trigger data upload:
curl --location --request POST 'http://localhost:8080/announcements'

Send a GET request to retrive data from ElasticSearch:
curl --location --request GET 'http://localhost:8080/announcements'
