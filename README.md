This project is REST Full API for managing currency rates, using docker with 2 container first currency-go container  for web app second is mysql container for database ,if you want to  deploy this app please follow this guide :

Quick Installation
===========
- download install file from https://raw.githubusercontent.com/viyancs/go-currency/master/install.sh
- $ chmod +x install.sh
- $ ./install.sh

Manual Installation
==========
- $ mkdir go-currency
- $ cd go-currency
- $ git clone https://github.com/viyancs/go-currency.git .
- $ cd src
- $ docker build -t viyancs/golang-mysql-api --no-cache=true . (making image)
- $ docker run -p 3306:3306 --name mysqldb -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=test_go -e MYSQL_ROOT_HOST=% -d mysql:5.7 (run mysql container)
- $ docker run -p 8282:8282 --link mysqldb:mysqldb --name currency-go -d viyancs/golang-mysql-api (run web app container that provide http REST API)
- $ docker exec -ti currency-go sh -c "cd cmd/schema && go build -a -o bintask task.go && ./bintask orm syncdb -force && ./bintask sample" (build and rebuild table with data sample)


API Docs Usage
==========
- GET http://localhost:8282/exchange (to get all exchange)

``` bash
{"code":200,"results":[{"avg 2 days ago":"90","c_from":"GBP","c_to":"INR","date":"2018-07-25 00:00:00","id":"4","rate":"90.4866"},{"avg 2 days ago":"14750","c_from":"USD","c_to":"IDR","date":"2018-07-25 00:00:00","id":"5","rate":"14750"},{"avg 2 days ago":"1","c_from":"USD","c_to":"GBP","date":"2018-07-25 00:00:00","id":"6","rate":"0.761913"},{"avg 2 days ago":"0","c_from":"USD","c_to":"JPN","date":"2018-07-25 00:00:00","id":"7","rate":"insufficient data"},{"avg 2 days ago":"0","c_from":"JPN","c_to":"GBP","date":"2018-07-25 00:00:00","id":"8","rate":"insufficient data"},{"avg 2 days ago":"130000000","c_from":"BTC","c_to":"IDR","date":"2018-07-25 00:00:00","id":"10","rate":"130000000"}],"msg":"","total":6}

```

- GET http://localhost:8282/exchange/(:date) (to get exchange by date example :date = "2018-7-25")

```bash
    {
    "code": 200,
    "results": null,
    "msg": "12",
    "total": 0
    }
```

- POST http://localhost:8282/exchange (insert new exchange) with body params like this :
    - date => '2017-08-25'
    - rate => '100000000'
    - from => 'BTC'
    - to   => 'IDR'
    

    
- POST http://localhost:8282/track (insert new track without exchang eand date)
    - from => 'WAVES'
    - to   => 'IDR'
    
- PUT http://localhost:8282/exchange (delete track by {from} and {to} field)
    - from => 'WAVES'
    - to => 'IDR'
    
 ``` bash 
 {
    "code": 200,
    "results": null,
    "msg": "success delete",
    "total": 0
}
 ```
