# Product Catalog API

## Dynamic reloading

App uses [gow](https://github.com/mitranim/gow) to reload project automaticaly

- `make install-binares` installs gow
- `make run` runs app using gow

[modd](https://github.com/cortesi/modd) is also an interesting alternative. It can be more general and more configurable than **gow**.

## Template libraries

App uses native default 'html/template' library. But there are alternatives like:

- [plush](https://github.com/gobuffalo/plush) - inspired by Ruby on Rails

## Templates

App uses 'embed' native package to build templates as file system (FS). So the application can read them from the binary directly.

## Database driver

App uses [pgx](https://github.com/jackc/pgx) driver to connect to Postgres.

## SQL instead ORM

- Each ORM is unique and needs to be learned. Once you know SQL, you can apply 99% of that knowledge anywhere.
- Once you learn an ORM you only know that specific ORM.
- It's easier to optimize when necessary.

## Authentication

App does not use third party authentication because the own authentication is cheaper and I can understand and practice security details.

- App uses [bcrypt](https://en.wikipedia.org/wiki/Bcrypt) hash function with salt to store password. The hash collisions are very unlikely.
- The salt is generate by bcrypt package and it's in hash result. Because of that, we can change the cost of hash function that the package will know how to compare hashes and passwords.

## Cross-Site Request Forgery (CSRF)

App generates and validate CSRF tokens using a HTTP middleware (https://github.com/gorilla/csrf) to protect against this kind of attack.

## Obfuscation (session)

Obfuscation is the process of making the data in a cookie unclear to attackers, making it almost impossible for them to determine how to generate valid data.

For instance, rather than storing a user’s email address in the cookie, we could instead generate a random string for each user:

We could then store the random string inside of our cookie. Now when a user makes a web request, we can look up who the user is via our table, but an attacker wouldn’t be able to generate a valid random string because they are just that - random strings. An attacker would need to guess random values hoping that one of them works.

App authentication system uses this approach to prevent cookie tempering.

### Why aren't we using JWTs?

JWTs are more complex to implement and don’t offer any major benefits here. Beside that, sessions is part of refresh tokens in JWTs.

### Generating session tokens

The native package "math/rand" generates pseudo random number based on a seed. It's not safe for the application because if attackers dicover the seed, they can generate valid tokens. App uses native package "crypto/rand" because it generates almost real random numbers.

The session token has 32 bytes (1x10⁷⁷ possibles) because of that the app can be a lot of users without conflict. And it increses the security, it ensures that it's virtually impossible for an attacker to guess a valid session token.

Read more: www.owasp.org/index.php/Insufficient_Session-ID_Length

### Sessions

The first app version has a sesstion table that allows only one session for each user.

## Migration

[goose](https://github.com/pressly/goose) handles the app migrations. It's a CLI installed using `make install-binaries`. If you're using asdf to manage go version, remember to reshim your go version after the installation.

`goose -dir migrations create {name_of_file} [sql|go]` -> create migration file
`goose -dir migrations postgres "{postgres_dns}" status` -> show migration process status
`goose -dir migrations postgres "{postgres_dns}" up` -> run pending migrations
`goose -dir migrations postgres "{postgres_dns}" down` -> rollback last migrations executed
`goose -dir migrations postgres "{postgres_dns}" reset` -> rollback all migrations executed


Pay attention to the order of migrations files. Working in groupssome migrations can be behind the current schema version. An approach to solve it is use sequencial numbers and not timestamp. `goose fix` rename files automatically.

### Goose as package

App uses goose CLI, but there is a package that allow generate a binary to run migrations inside the code. `go run cmd/migration/main.go`

## Context

App uses context to set, get and require user data. There is a middleware to handle this. That way, we can know request's user at any time by getting it from context.

## Env

App uses [godotenv](github.com/joho/godotenv) to load envs from .env file

## Email

App can send email to users

### Reset password

Generate token and url to return to website and reset password. These token is verified only once, so the same token can't reuse

### Smtp

App uses [Mailtrap](https://mailtrap.io) to send reset password email.

## Errors

App separetes public errors which are erros that can be shared with users; and internal erros which are errors to log and trace application.

## Other securities apects (not implemented in this software)

- Block ip that make a lot of invalid requests
