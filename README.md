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