# bcrypt stuff...


run bcrypt to see if hashes match

```bash
go run cmd/bcrypt/bcrypt.go compare \
"fake123" \
'$2a$10$qUz2nP49GyaFF85GgzHGMuYB4YPgVT0Hv395q..jhNTm54VCqtnSe'
Password is correct
```

run bcrypt to see if hashes don't match

```bash
go run cmd/bcrypt/bcrypt.go compare \
"fakeABC123" \
'$2a$10$qUz2nP49GyaFF85GgzHGMuYB4YPgVT0Hv395q..jhNTm54VCqtnSe'
Invalid password fakeABC123
```