# Wildfire Backend Assessment
An app that fetches a random name and joke, and combines them.

## APIs Used

  - https://names.mcquay.me/api/v0
  - http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=nerdy
## Technologies Used

  - Go
  - HTTPie


  ## Features
 1. User can view a random joke

## Preview
 <img src="./assets/wildfire-preview.gif">

 ## Development

 ### Getting Started
 1. Clone the repository
    ``` bash
    git clone https://github.com/kusholhuq/wildfire-backend-assessment.git
    cd wildfire-backend-assessment
    ```
 1. Once in the project directory, run the main.go file to start the server
    ``` bash
    go run main.go
    ```
  1. Open another terminal window and make a GET request to port 5000 (Below HTTPie is being used, but Postman, the browser or any other service also works)

    http --body localhost:5000
