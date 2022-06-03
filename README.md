# Ta-vivo audit log

This is a service to save log history like check updates, integration test requests and much more of Ta-vivo APP, written in Golang.

## Get started

Install packages
```
go get
go install
```

Run 

```bash
go run ./src/
``` 

## Example of JWT creation

The service work with JWT, you can generate a JWT token with the following command:

```bash
const jwt = require('jsonwebtoken');

const token = jwt.sign(
  { name: 'ta-vivo-api' },
  'super_secret',
  {
    expiresIn: '1y',
  }
);

console.log(token);
```