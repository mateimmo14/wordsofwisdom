# Words of Wisdom
A small website where you can share funny quotes
## Try it out!
Try it out at <a href="https://wordsofwisdom.duckdns.org">wordsofwisdom.duckdns.org</a>
## Setup
If you want to try it out for yourself:
* Clone the repository
* Setup the environment: `export DATABASE_URL=<your database url here>`
* Setup the database, in the sql/schemes directory run `goose postgres $DATABASE_URL up`
* Build the binary, in the root of the repo run `go build`
* Run it with `./wordsofwisdom`
## Requirements
* Go 1.26
* PostgreSQL v15+
* goose
