# collatzExercice

## Requirement

- Golang
- docker
- docker-compose
- minikube


## Setup

Make sure minikube is running : `minikube start`

`make deploy` - this will setup code in minikube

## Find URL

`make getURL` - this will return the url to run the code

```
$ make getURL
minikube service collatz-app-service --url
http://192.168.64.4:30152
```

## Run

URL: http://minikubeIP:port/collatz/{value}

In browser 
     http://192.168.64.4:30152/collatz/21
     
Terminal
    ` curl http://192.168.64.4:30152/collatz/1`
    
## Run script

`make runScript n="22"` - run code from 1 ... n  


## Testing

`make test` - this will run some test

