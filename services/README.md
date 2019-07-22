# services

This folder contains example services that work together with the CLI app to create a simplistic demo system. The basic idea is there are **users** who can create **orders** for **products**. 

## running the services

>Before starting the services, be sure to copy the `sample.env` file to just `.env` in the root of the repository. Then make any changes needed for your environement.

You can start the services by typing the following commands in root of this repository:

```
> go run sercices/users/cmd/main.go
```

```
> go run sercices/products/cmd/main.go
```

```
> go run sercices/orders/cmd/main.go
```

*Each service requires it's own terminal/shell.*
