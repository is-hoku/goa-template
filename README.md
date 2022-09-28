# goa-sample
This is a sample REST API using [Goa v3](https://github.com/goadesign/goa/tree/v3).

## Usage
### Generate presentation layer code from design
After the design is done, run this in the `/webapi` directory to create files. Note that `make goagen` removes all client files generated by `goa gen`.  
```
make goagen
```
### Build and Run
Basically, [air](https://github.com/cosmtrek/air) builds the code automatically, but if you want to build it manually, run this.  
```
make build
```
And you can execute the binary.  
```
make exec
```
### Generate infrastructure layer code from SQL
After writing queries in SQL, run this to generate code. For more information, you can see the [sqlc](https://github.com/kyleconroy/sqlc) repository.
```
make sqlc-generate
```
### Migrate
After writing [atlas](https://github.com/ariga/atlas) ddl, run this to execute the migration.
```
make migrate
```
You must generate SQL files of the table definition and run the sqlc command.  
Run this and everything will be fine.
```
make auto
```
### Generate class diagram text
Run this in the app container to generate class diagram text compatible with plantuml.
```
make plantuml
```
