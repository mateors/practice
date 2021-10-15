# Sqlite3 database command practice
SQLite is the most used database engine in the world. \
SQLite is built into all mobile phones and most computers and comes bundled inside countless other applications that people use every day

> The complete state of an SQLite database is usually contained in a single file on disk called the "main database file".

## Check the version of the sqlite db
> file imapsql.db
``imapsql.db: SQLite 3.x database, last written using SQLite version 3031001``

## CLI
> apt install sqlite3

## Important commands
> **.show**  how the current values for various settings \
> .tables ?TABLE? \
> .databases \
> .dbinfo ?DB? 

> .output backup.txt \
> **.dump** Render all database content as SQL into backup.txt file \

> .schema ?PATTERN? \
> **.quit**   Exit this program

## Resource
* https://www.sqlite.org/cli.html
* https://sqlitebrowser.org
* github.com/mattn/go-sqlite3