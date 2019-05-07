# redeam
### test task

#### Description
    Using Go as your language, create a CRUD API to manage a list of Books, fulfilling the following requirements:
    
    
    1. Books should have the following Attributes:
    
    
    - Title
    
    - Author
    
    - Publisher
    
    - Publish Date
    
    - Rating (1-3)
    
    - Status (CheckedIn, CheckedOut)
    
    
    2. Each endpoint should have test coverage of both successful and failed (due to user error) requests.
    
    
    3. Use a data store of your choice.
    
    
    4. Include unit tests that exercise the endpoints such that both 200-level and 400-level responses are induced.
    
    
    5. The app should be stood up in Docker, and client code (such as cURL requests and your unit tests) should be executed on the host machine, against the containerized app.
    
    
    6. Send the project along as a git repository.
    
    Again, I want to be sensitive to his time, so I would suggest no more than 8 hours work on this, but it's crucial for us to get a sense of his coding skills!
    
#### Init app
To init app you need move the make commands:

    make init
    cd docker
    docker-compose up -d

#### To test app
    go test -race ./... 
#### Linter
I use [golangci-lint](https://github.com/golangci/golangci-lint)
##### Linter for IDE
    https://www.jetbrains.com/help/go/settings-tools-file-watchers.html

#### Postman link to collections
[Books Api Collection](https://www.getpostman.com/collections/64a2a9df0e22949de401)

