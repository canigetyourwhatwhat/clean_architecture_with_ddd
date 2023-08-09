# clean_architecture_with_ddd
It is dedicated to demonstrate the proposed approach: combination of Clean Architecture and DDD

## Prerequisite:
-  You have installed Go version 1.20
-  You have installed MySQL locally or you run MySQL docker container.

## Directory structure
- cmd (it contains files to start the application.)
    - cmd.go
        -  it builds the server.
    -  router.go
      -  it holds available endpoints.
- config
    - it defines variables for config.yml.
- database
    - it sets up database such as creation of table, inserting default records.
- internal
    - it holds all the business domain
        - **controller (API related)**
            - api
                - it writes about endpoint
            - entity
              - request
                - request from the API
              - response
                - response for the API
            - inputPort
              - validate API inputs 
            - middleware
              - middleware for the server
            - presenter
              - it converts models to the required models for response
        - **entity**
          - Whole models for business logic
        - **infrastructure**
          - this directory holds the implementation of interface directory
          - repository
        - **interface**
          - this has interface for all external devices
          - repository
        - **usecase**
          - it holds all the main business logic


## To run the server
```go
go run main.go
```