FROM golang:1.22.2-alpine
RUN mkdir app
# run mkdir app = creates a folder named build in the working directory or creates a working directory
WORKDIR /app  
# workdır /app = sets the working directory of the container specified in dockerfile sets the working directory to app
COPY . .
# The first dot in the dockerfile refers to all the files found in the dockerfile and handles them 
# the second dot represents where to send, i.e. app 'i'
RUN go build -o ascii-art-dockerize . 
# this is where the build is done - the part named -o is ascii... named go build if = build command
ENTRYPOINT [ "./ascii-art-dockerize" ]
# specifies the entrypoınt to run . = specifies the working directory. asciiart... = file to be executed.
EXPOSE 8080
# portal to work        