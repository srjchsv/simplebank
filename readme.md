### 🏦 Educational project- simplebank🏦

## Currently made:
- Database schema design. 
- Generated CRUD queries using SQLC.
- Implemented balance transfer transaction.
- TX lock handled.
- Unit and integration tests.
- Github actions CI
- REST api. Currently available:

   <em>First run the authorization mircoservice https://github.com/srjchsv/service
to be able to signup and signin, and use all other features.</em>

![simplebank schema](static/simple-bank-schema.jpeg "simplebank database schema")


### Accounts management:
    - POST signup `make signup`
    - POST signin `make signin`
    - GET get account by id `make get-account`
    - PUT Update account by id  `make update-account`
    - GET get accounts in batches `make get-accounts`
    - DELETE delete account by id `make delete-account`

### Balance transfers management:
    - POST Transfer balance from-to account `make transfer`

## To run server use `make server`

### db migrations `make up`

🚧 More features to be made... 🚧

### Below is the postgres database schema:

![simplebank database schema](static/simplebank.png "simplebank database schema")


### Useful links:
-  videos:
https://www.youtube.com/watch?v=rx6CPDK_5mU&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE

- articles:
https://dev.to/techschoolguru/series/7172