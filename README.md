# Xavier
Xavier is the API server for Arya. Many awesome

## Requirements
- Go 1.4.x
- PostgresSQL 9.4.x
- Ruby 2.2.x with Rails 4.4.x

## Installation on your local machine
1. Make you have all the requirements update and running. This means installations of the latest Go (GOPATH setup properly), PostgreSQL and Ruby on Rails (make sure you can use Bundler and Rake). Homebrew helps a lot on OS X. 
2. Start PostgreSQL. 
3. Move the the root of project (folder 'xavier')
4. (Optional) If you haven't setup the database yet, run the following commands. 
  - createdb "arya_development"
  - rake db:migrate db=development (sets up the database tables)
  - rake db:seeds db=development (puts in some dummy data)
5. (Optional) Setup all the Go dependencies via the Makefile. Run 'make install' and you're set!
6. Run 'fresh' and the project will be build. Also it will be recompiled after you changed and saved a file. Woot!
7. (Optional) Install pgweb to easily look inside the PostgreSQL databases https://github.com/sosedoff/pgweb


##TODO List
- [ ] Registration is still missing. (with confirm mail)
- [ ] Localization support. For questions just make a new record in the db and refer to this record.
- [ ] Let people send invites to connect to each other (therapist->patient viceversa).
- [ ] Receive related users to an user.
- [ ] Let an related access all information available based on permissions.
- [ ] Storage unit tests, mock storage tests.
- [ ] Storage DB pagination support (one page 1 fetch 20 items, on page 2 the next 20 items)
- [ ] Send push notification(s) via Goroutine.