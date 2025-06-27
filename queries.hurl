POST {{HOST}}/shorten
{
    "long_url": "https://google.com"
}
HTTP 201

GET {{HOST}}/redirect/abc
HTTP 200