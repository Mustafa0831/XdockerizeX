# ASCII-ART-WEB-DOCKERIZE

## Objectives

The main goal of the subproject _**DOCKERIZE**_ is to create a Dockerfile (== an executable that contains the files and the dependencies of a program.), one image and one container

## How to do a docker ?

You have to create a Dockerfile which contains the following :

# The base go-image
FROM golang:1.16.3

# Create a directory for the app
RUN mkdir /ascii-art-web-dockerize
 
# Copy all files from the current directory to the app directory
COPY . /ascii-art-web-dockerize
 
# Set working directory
WORKDIR /ascii-art-web-dockerize
 
# Run command as described:
# go build will build an executable file named main in the current directory
RUN go build -o main . 
 
# Run the main executable
CMD [ "/ascii-art-web-dockerize/main" ]
```


Afterwards, in order to run the docker in the command prompt:

You have to build the program with the command :

```cmd
docker build -t ascii-art-web-dockerize .
```

Then, run it with the command :

```cmd
docker run -it --rm -p 8000:8001 ascii-art-web-dockerize
```


<hr>

# What we have learned from this Project

This project helped us learn about the topics below : 
- Client utilities.
- The basics of web :
    - Server
    - HTTP 
    - HTML
- Learning about docker.
- Using and setting up Docker :
    - Services and dependencies.
    - Containerizing an application.
    - Compatibility/Dependency.
    - Creating images.
- Ways to receive data.
- Ways to output data.

<hr>

# This project was made by :
### Mustafa (as mus11110)
### Zhangir (as zhangir11)
