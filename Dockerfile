#tart from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/DanielDanteDosSantosViana/darth_vader

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/DanielDanteDosSantosViana/darth_vader

#ENV
ENV ID_AWS="testewas"
ENV SECRET_KEY="secretawskey" 
# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/darth_vader

EXPOSE 3000