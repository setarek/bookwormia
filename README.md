# Bookwormia

This project is composed of three distinct services. The "user" service manages user authentication. The "admin" service facilitates the seamless addition of single or bulk books into a Redis queue, which is then consumed by another service, ultimately writing the information into the database. The "book" service operates efficiently with multiple workers, extracting newly added or modified book details from Redis and persisting them to PostgreSQL. Furthermore, this service exposes APIs for accessing book lists, retrieving detailed book information, and facilitating user actions such as submitting reviews or bookmarking a book.

## How to run project

### Run manually

Ensure that your system includes both Postgres and Redis.

* login to Postgres CLI

  ```bash
  sudo -u postgres psql
  ```

* Create database

  ```sql
  create database test;
  ```

* Create user and password

  ```sql
  create user test with encrypted password 'test';
  grant all privileges on database test to test;
  ```

* This project is implemented using the Go workspace. Before running any service, ensure to execute the following command to download the project and its dependencies:

  ```bash
  go work init
  go work use user
  go work use admin
  go work use book
  go work use pkg
  ```

  Run (or build then execute) each service separately

* ```
  go run user/main.go
  go run admin/main.go
  go run book/main.go
  ```



* Add one book

  ```
  go run admin/main.go addNewBook --name "Harry Potter" --details "Lorem ipsum"
  ```

* Add multiple book with csv file

  ```
  go run admin/main.go newBooks --name "books.csv" --path "your path"
  ```

  csv file example:

  ```
  Name,Description
  Think Python,test1
  100 mistake in olang,test2
  ```

  

* Edit one book

  ```
  go run admin main.go editBook --id 2812 --name "Test" --details "Test"
  ```

* Edit multiple books

  ```
  go run admin/main.go editBooks --name "books.csv" --path "your path"
  ```

  csv file example:

  ```
  Id,Name,Description
  55,Think Python,test1
  66,100 mistake in olang,test2
  ```

  

* delete one book

  ```
  go run admin/main.go deleteBook --id 2812
  ```

* Delete multiple books

  ```
   go run admin/main.go deleteBooks --ids "2809,2808"
  ```

  

