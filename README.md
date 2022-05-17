# GO Backend

---
This projects was made in Go Using ORM (Gorm) and fiber, you can see the documentation here:
* [Fiber](https://docs.gofiber.io/)
* [Gorm](https://gorm.io/index.html)

#### Steps:

1. Install Go
2. Create a DataBase, in this case this example runs using Postsgresql
3. Look for the file _.env_ in this project and change the values

    ```
        DB_HOST= localhost
        DB_NAME= # Your db name
        DB_USER= # your db user
        DB_PASSWORD= # your db password
        DB_PORT= 5432 # this is a default port you can change the port
        PORT_LISTEN = 3000 # the port listen in this case is 3000
    ```
4. Run the command _go get_ for update the packages
5. Run the command _go run main.go_ for run the project 

#### Documentation:

the path for run the instructions are next:

* _<your localhost:port >/api/user/_
* _<your localhost:port >/api/address/_

The routes were declared in the folder routes, each folder represent the route 