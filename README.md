# go_backend

this project is being developed for the backend course @pucpr

## project structure
>
> ```server.go``` - the main entry point
> > ```app/handlers``` - directory with route handlers
>

## stack used

in this project, we're aiming to get the best possible performance, not giving up on DX
> echo - our back-end framework, based on the net/http standard lib
>
> air - hot reload for our application

## requirements for the api

- [ ] SQLite integration
- [ ] Swagger
- [+ -] Testing (kinda....)
- [ ] Logging
- [ ] Authentication / middleware
  - [ ] JWT
  - [ ] Roles
    > Middleware already kinda implemented, check the ```server.go``` file for the implementations
- [x] Routing

## todo / next steps

- [ ] add propper logging (zap!)
- [ ] add middleware for auth
- [ ] Add the ```sqlite3``` support (i.e. fix the gcc error)
- [ ] add user role changing by file (.config) (??)

- [x] make ```database.go``` acc file with dbs
- [x] create post handlers and operations
