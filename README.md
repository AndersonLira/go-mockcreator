# Go - Mockcreator
A basic SOAP mock creator that uses origin soap service

Mockcreator was developed in Golang and it is a ASIS of MockCreator Java project.

This project uses https://github.com/golang/dep to manager dependencies. To build:
```sh
dep ensure
go build
```
### Important

Some variables should be defined on your environment.

  - SERVICE_URL= origin url service
  - MC_USER= origin service user name
  - MC_PASS= origin service password
  - SERVER_CONTEXT = context of service example mockservice
 

You also need to create a folder named **payloads** where executable is. Is necessary **config.json** present at same executable folder

## Configuration

The file of configurations is explained: Each key changes mock server behaviour.
- **returnDelay** is used to delay return of methods configured in **delayMethods** property.
- **showErrorServer** if true service fault return will be logged.
- **workingAsProxy** if true mocks will be ignored. All requests will be in origin server.
- **clearCache** methods that need remove cache after some other method call.
- **cacheEvict** methods that should not be cached. These requests will be always in origin server.
- **staticReturn** methods that should have these returns statics. Here is configured the file of return.
- **logRequestBody** It will print request soaps body.
- **logResponseBody** It will print response soaps body.

To see config alterations is necessary restart application

