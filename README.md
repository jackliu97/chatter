### Chatter

#### Setup
1. Run `./bin/build && ./bin/startsql`
2. Take the ip output in the final line and update config.yaml's `db_ip` with that value.
3. Run `./bin/startapp`
4. The app should be visible on `http://localhost:8080/`


 - you can run `./bin/stop` anytime to stop and remove containers
 NOTE: There's no volume binding, so you will lose ALL of your chats + users.
 
 #### Usage
  1. Create a user by providing a unique username and a password, you should be logged in on success.
  2. Start chatting
 
 
 #### Endpoints
 
  - Create users
    
    /user `POST`
    
    Payload
     - username
     - password
     
    Example
    ```
    Content-Type: application/json
    
    {username: "jackliu", password: "jack!123"}
    ```
    
    Success: A HTTP 201 response
    
    Failure: A HTTP 400 response


  - Login user
  
    /login `POST`
    
    Payload
         - username
         - password
         
    Example
    ```
    Content-Type: application/json
        
    {username: "jackliu", password: "jack!123"}
    ```
    
    Success: HTTP 200 response
    
    Failure: HTTP 401 response
    
  - Get Messages
    
    /messages `GET`
    
    Example
    ```
    /messages?page=1&size=15
    ```
    
    Sample response
    ```
    Content-Type: application/json; charset=utf-8
    
    {
        "code": 200,
        "data": [
            {
                "username": "jack",
                "message": "a"
            },
            {
                "username": "jack",
                "message": "asdf"
            },
            {
                "username": "jack",
                "message": "asdf"
            },
            {
                "username": "john",
                "message": "asdfasdf"
            }
        ],
        "error": ""
    }
    ```
    
    Notes: 
     - First page starts at 1
     - Invalid page will default to 1
     - Invalid size will default to 10
    
    
    
    
