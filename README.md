
# ecommerce

This project exposes a post endpoint `/checkout` which accepts a list of product IDs to calculate the total price of all the items.

## Requirements

Download >= go1.19

## Solution

For a straight forward and simple implementation, `items.json` acts as a datastore. All the products are read from the file and loaded onto the application while starting the server.
A idiomatic flow of code consisting of controller which manages http codes & validations, service code which manages calculation of total price.

Run `make build -> make run`

## Future Actions

1. Improve request validation and clear requirements for invalid/error scenarios - should a request be processed with invalid ids? should we generate partial checkout response of only valid ids? should we also add the invalid ids in the checkout response?
2. Setting up a database(mySQL) to replace the `items.json` file.
3. Use of more idiomatic solutions for reading config files. Viper - https://github.com/spf13/viper
4. Dockerization of the database and application.
5. Admin tool to manage the items repository manually.


## Commands

Build 

`make build`

Test

`make test`

Run

`make run`


