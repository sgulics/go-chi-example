# Golang Chi Web Application POC

Investigating using [Chi](https://github.com/go-chi/chi) to create a simple web application that supports the following:
1. REST API
2. Admin Front End
    1. HTML Templates with ability to change template without restarting
    1. FuncMaps 
    1. Javascript/CSS with hot reloading
    1. Asset support (images, css, javascript)
    1. Login and Session Management
    1. Support for Flash messages
    
## Build & Run

`go run cmd/server/main.go`

Server is running on http://localhost:3333

## Routing

### API

* GET http://localhost:3333/v1/articles - Get List of articles
* GET http://localhost:3333/v1/articles/1 - Get Article by ID
* PUT http://localhost:3333/v1/articles/1 - Update Article
* DELETE http://localhost:3333/v1/articles/1 - Delete Article


### Monitor 
* GET http://localhost:3333/monitors/ping - Simple route to ping the API


### Admin

In order to use the Admin front end you need to install webpack:

`npm install`

for development mode (start and leave running in another console)
`./node_modules/.bin/webpack-dev-server --config webpack.config.js --hot --inline`

Or for production mode
`./node_modules/.bin/webpack --config webpack.config.js --bail`

All assets are in the assets folder

All templates are in the templates

Please note the admin front end does not actually do anything useful, I just wanted get an understanding of how all of this works.

Go to http://localhost:3333/admin 

When prompted to login, to in anything.

You should be can make changes to the HTML templates without restarting. 

You can change the CSS and Javascript and if in development mode, the assets should hot reload

 


    
  
 
  


