# Take Home Test


## Problem Statement
You are asked to develop an API to simulate transactions flow


## Services
### User Services
* `register` - Creates the customer if not exist
* `login` - Logs in as this customer
* `logout` - Logs out of the current customer
* `balance` - Show User balance of the customer

### Transactions Services
* `topup` - Add this amount to the logged in customer balance
* `withdraw` - Withdraws this amount from the logged in customer
* `transfer` - Transfers this amount from the logged in customer to the target customer
* `history transaction` - History transaction from the logged in customer


## Expected
* `Register`
   Response :
    Registration success
    Hello, Alice!
* `login` 
   Response :
    Hello, Alice!
    Your balance is $0
* `balance` 
   Response :
    Hello, Alice!
    Your balance is $0
* `logout`
   Response :
    Goodbye, Alice!

* `topup`
   Response :
    Topup success!
    Your balance is *any amount that have submitted*
* `withdraw`
   Response :
    Withdraw success!
    Your balance is *any amount*
* `transfer`
   Response :
    Withdraw success!
    Your balance is *any amount*

* `history transaction`
   Response :


## Criteria
* Using docker as container
* Using Golang framework, feel free to use any framework (gin, echo, etc.)
* Preferably apply the design pattern
* Preferably apply any authentication (token, jwt, Auth User, etc)
* You can use any database (SQL, NoSQL, etc)
* Unit Testing is a plus point


## Collect
* Please submit the test no more than 3 days after receiving this test
* Please zip the test when you have done or you can send the git URL
* Dont forget to give the instructions to run your code using Readme.md 
