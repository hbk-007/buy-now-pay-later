# Buy-Now-Pay-Later Service

Description
This is a command line application for buy now paylater service. 

## Features
1. User can registered with a specified credit Limit
2. Merchants can registered
3. Various types of reports can be generated.
    1. how much discount we received from a merchant till date
    2. dues for a user so far
    3. which users have reached their credit limit
    4. total dues from all users together

## This project has 4 Domain layer :
1. Entity Layer
2. Repository Layer
3. Usecase Layer
4. Controller/Delivery Layer

Layers are implemented as interfaces thus easier to incorporate changes. 
For e.g. the repository layer can easily plug mysql/ mongo instead of my in-memory database representation. 

## How to run this project
```
$ go run .
```
