# http_ad_testing
 small Go program to make HTTP requests for Ad testing

config.json looks something like:

{
"prefix": "URL PREFIX",
"suffix": "URL SUFFIX"
}

mac_list is just a list of MAC addresses or device IDs or whatever you need to mix into the URL

main.go will mix the devices into the URL and make a request to that full URL