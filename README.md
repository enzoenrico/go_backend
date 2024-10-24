# go_backend
this project is being developed for the backend course @pucpr

## project structure
>
> server.go - the main entry point
> > app/handlers - directory with route handlers
>

## stack used

in this project, we're aiming to get the best possible performance, not giving up on DX
> echo - our back-end framework, based on the net/http standard lib
>
> air - hot reload for our application

## requirements for the api

- [ ] SQLite integration
- [ ] Swagger
- [ ] Testing
- [ ] Logging
- [ ] Authentication / middleware
- [x] Routing

## todo

- [ ] Add the sqlite3 support (i.e. fix the gcc error)
- [ ] add middleware for auth
- [ ] add user role changing by file (.config) (??)
- [ ] add propper logging (zap!)

- [x] make database.go acc file with dbs
- [x] create post handlers and operations
