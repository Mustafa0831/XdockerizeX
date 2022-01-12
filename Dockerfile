# The base go-image
FROM golang:1.16.3
LABEL name="ASCII-ART-WEB DOCKERIZE"
LABEL description="Docker"
LABEL authors="mus11110; zhangir11;"
LABEL release-date="17.09.2021"
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

#Run the docker in the command prompt:
#1
#Build the program
#docker build -t application-tag .

#2
#Run the docker
#docker