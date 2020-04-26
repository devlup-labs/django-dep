# django-dep  

[![GoDoc](https://godoc.org/github.com/devlup-labs/django-dep?status.svg)](https://godoc.org/github.com/devlup-labs/django-dep)
[![Go Report Card](https://goreportcard.com/badge/github.com/devlup-labs/django-dep)](https://goreportcard.com/report/github.com/devlup-labs/django-dep)
[![GolangCI](https://golangci.com/badges/github.com/devlup-labs/django-dep.svg)](https://golangci.com)
[![GitHub Release](https://img.shields.io/github/release/devlup-labs/django-dep.svg?style=flat)](https://github.com/devlup-labs/django-dep/releases)

## Introduction  
This utility is aimed to trigger deploys on VM that are behind a firewall. This utility can run behind a proxy server and execute a script on the host machine to re-deploy the applciation.  

## Installation  

The server can be started with
```shell
make start
```  

To install dependencies
> Note: start and compile steps already install dependencies first  
```shell
make install
```  

A binary can be generated with
```shell
make compile
```  
