# MBE TakeHome Test
## Problem Statement
Develop an API to simulate transactions flow

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

## Database
* Using postgresql, setting on .env file

## Go framework
* Using Gin