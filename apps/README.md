# apps

This folder is for applications that involve user interaction. In the **cli** directory you'll find a very basic command line interface for placing orders. The app depends on the **users**, **orders**, and **products** services found in the **services** folder at the root of this repository.

## running the cli app

>Before starting the services or the CLI app, be sure to copy the `sample.env` file to just `.env` in the root of the repository. Then make any changes needed for your environement.

After you have started the dependent services, you can start the CLI app by typing the following command in root of this repository:
```
> go run apps/cli/cmd/main.go
```
