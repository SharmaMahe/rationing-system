# Martian Rationing System
You are part of the Ares III mission to Mars exploring “Acidalia Planitia” plain on Mars in the year 2035. Due to unfortunate circumstances involving a dust storm, you got stranded on Mars alone with no communication to your team or Earth’s command center.
Your surface habitat ("Hab") on Mars contains a limited inventory of food supplies. Each of these supplies is in the form of a packet containing either Food or Water. The food items have details of the content, expiry date and calories they provide. The water packet contains only the details of the quantity in liters and doesn’t expire.
To survive, an average human needs 2500 calories and 2 liters of water per day. You are also worried about the wastage that might occur if the food is not consumed before the expiry date. Being a brilliant programmer, you have decided to create an algorithm that can ration the food based on daily requirement and provide you inputs on which packet to consume on a given day and how many days will you survive with the available inventory.

You can find more info here - https://docs.google.com/document/d/12iUPtQJdN5tspzy2jSz8bqUytWMhcua3Cz38kGXq_Mo/edit?usp=sharing

## Prerequisites
    'go' should be installed on your machine.

## Installation
	Execute following commands :
	1 Close the repo and go to src/app
    2 Run "go get github.com/astaxie/beego"
	3 Run "go get github.com/go-sql-driver/mysql"
	4 Run "go get github.com/beego/bee"
    5 After that set the GOPATH,GOROOT and PATH variables 
    6 Configure the database credential in src/app/conf
    
## Usage

	Run "bee run" it will start server on localhost:8080 (if you are running system on local machine)


## Test Case
	To execute test cases run 
	go test test/default_test.go