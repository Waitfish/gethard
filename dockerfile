FROM golang
  # We create an /app directory within our
  # image that will hold our application source
  # files
RUN mkdir /app
  # We copy everything in the root directory
  # into our /app directory
ADD gethard /app
  # We specify that we now wish to execute
  # any further commands inside our /app
  # directory
WORKDIR /app
  # we run go build to compile the binary
  # executable of our Go program
# RUN go get github.com/imroc/req && go get github.com/shirou/gopsutil
#RUN go build -o gethard .
  # Our start command which kicks off
  # our newly created binary executable
CMD ["/app/gethard"]
