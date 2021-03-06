# simple-golang-webapp
A webapp made for a devops / cloud deployment workshop run by Slalom & AWS, intended to be an example application that can be deployed to cloud infrastructure and automatically configured using Ansible.

<img width="654" src="https://raw.githubusercontent.com/mattchilds1/simple-golang-webapp/master/screenshots/screenshot.png" />

## Config

This simple app requires one environment variable for the hostname/ip of the mongodb server.

```bash
DB_URL=localhost
```

Also ensure go is properly installed and configured (e.g. $GOPATH etc) on your machine.

## Usage

### Database
You will need to have mongodb running on your machine and have imported the database/collection in reports.dataset.json in this repository.

With mongodb running, import the datset with the following command:

```bash
mongoimport -d reports2 -c report --file reports-dataset.json
```

### Application
1. Get dependencies
```Golang
go get
```

2. Build binary
```Golang
go build
```

3. The previous command should produce a binary called 'simple-golang-webapp', you will be able to run this with the following command and access the application in your browser on localhost:8080
```bash
./simple-golang-webapp
```


## Related

Terraform & Ansible repository for deploying this application Amazon Web Services
https://github.com/nmarchini/tf-01
