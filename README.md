# Authlify - User mangament for APIs

Authlify is a small open-source API that can act as a self-standing API service
for handling user registration and authentication for JAM projects.

It's based on OAuth2 and JWT and will handle user signup and authentication.

It exposes the following endpoints:

* POST /signup

  Register a new user with an email and password.

  ```json
  {
    "email": "email@example.com",
    "password": "secret"
  }
  ```

  Returns:

  ```json
  {
    "id":1,
    "email":"email@example.com",
    "confirmation_sent_at": "2016-05-15T20:49:40.882805774-07:00",
    "created_at": "2016-05-15T19:53:12.368652374-07:00",
    "updated_at": "2016-05-15T19:53:12.368652374-07:00"
  }
  ```

* POST /verify

  Verify a registration or a password recovery. Type can be `signup` or `recover`
  and the `token` is a token returned from either `/signup` or `/recover`.

  ```json
  {
    "type": "signup",
    "token": "confirmation-code-delivered-in-email"
  }
  ```

  Returns:

  ```json
  {
    "access_token": "jwt-token-representing-the-user",
    "token_type": "bearer",
    "expires_in": 3600,
    "refresh_token": "a-refresh-token"
  }
  ```

* POST /recover

  Password recovery. Will deliver a password recovery mail to the user based on
  email address.

  ```json
  {
    "email": "email@example.com"
  }
  ```

  Returns:

  ```json
  {}
  ```

* POST /token

  This is an OAuth2 endpoint that currently implements
  the password and refresh_token grant types

  ```
  grant_type=password&username=email@example.com&password=secret
  ```

  or

  ```
  grant_type=refresh_token&refresh_token=my-refresh-token
  ```

  Returns:

  ```json
  {
    "access_token": "jwt-token-representing-the-user",
    "token_type": "bearer",
    "expires_in": 3600,
    "refresh_token": "a-refresh-token"
  }
  ```

## TODO

* Support for postgres
* Custom user data

## License

The MIT License (MIT)
Copyright (c) 2016 Mathias Biilmann Christensen

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.