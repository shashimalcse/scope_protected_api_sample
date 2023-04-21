# Scope Protected Mock Tweeter APIs

This repository contains mock Tweeter APIs that simulate the functionality of retrieving a list of tweets, getting a tweet by ID, creating a tweet, and deleting a tweet.

## API Endpoints

### GET /tweets
Returns a list of tweets.

Required Scope(s)
* tweet.read
* users.read

### GET /tweets/:id
Returns a tweet by its ID.

Required Scope(s)
* tweet.read
* users.read

### POST /tweets
Creates a new tweet.

Required Scope(s)
* tweet.read
* tweet.write
* users.read

### DELETE /tweets/:id
Deletes a tweet by its ID.

Required Scope(s)
* tweet.read
* tweet.write
* users.read

## How to Use
1. Clone this repository.
2. Install dependencies by running `go mod download`.
3. Start the server by running `go run cmd/server/main.go -config <config location>`.
4. Use a tool like Postman to make requests to the API endpoints.
