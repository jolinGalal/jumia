# jumia
It is a single page application for rendering customer data with ```sqlite``` database and ```golang``` programming language.<br />
 * not using build in function to validate phone number
 * filter by country and state 
 * order by CustomerID, CustomerName and CustomerPhone ascending or descending
 * add paging
 * use ``` goa ``` framework for creating rest API for more details about ```goa``` https://goa.design/learn/getting-started/
 * use ```design pattern, ddd and microservices architecture```
 * add unit test, Dockerfile and docker compose.yml

# Folder structur
```
.
├── build
│   └── customer
│        └── Dockerfile
│   └── docker compose.yml
├── cmd
│   └── customer
├── internal
│   └── customer
│   └── model
├── pkg
│   └── database
│   └── utils
└── README.md
```
# run project

 * navigate to ```build/docker compose.yml``` <br />
   run ``` docker-compose up ``` in the termenal
 * navigate to ```cmd/customer compose.yml``` <br />
  run ```go run *.go ``` in the termenal
  * use postman collection
     https://www.getpostman.com/collections/c0425fc81e9b8bd90ea7
# code description
* validate phone number by using ``` validation package ```internal/model/country``` check for example
```
var phone = "write your phone number"
var v = validation.New(&phone)
var cameronRegularExpression = func() {
	v.Code("(237)")
	v.Optional(' ')
	v.OneOf('2', '3', '6', '8')
	v.Digits(7, 8)
}
 v.Valid(cameronRegularExpression)
```
* use ```goa``` framework for building api check ```internal/customer/design``` 
```
dsl.Method("list", func() {
   dsl.Description("list customers")
....
})
```
* use ```database package ``` for creating sqlite database connection 
* use ``` utils package ``` for paging

