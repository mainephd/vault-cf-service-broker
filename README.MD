# Vault Cloud Foundry Service Broker
[![Build Status](https://travis-ci.org/mainephd/vault-cf-service-broker.svg?branch=master)](https://travis-ci.org/mainephd/vault-cf-service-broker)

# Getting Started
* Install Go for local development from [https://golang.org/dl/](https://golang.org/dl/). Also available from the same location are instructions on how to setup your development environment for go.

* Godeps for vendoring application dependencies will need to be installed
```
go get github.com/tools/godep
```
    
    Periodically you shoudl run `godep save` to update the contents of the `Godeps` folder with the current set of dependencies.

* Ginkgo is used for tests and needs to be installed seperately as they won't vendored with the repository.
```
go get github.com/onsi/ginkgo/ginkgo
```

    Running tests are accomplished by running `ginkgo -r` from the root directory of the project. 
    If you want tests to continuously run while files are being changed then run `ginkgo watch -r` from the root of the project.
    
# Builds
Builds of this project can be done locally with `go install` and on each commit the project is compiled and tested via 