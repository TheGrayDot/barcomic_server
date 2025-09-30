FROM golang:1.24-bookworm

WORKDIR /barcomic
COPY ./cmd /barcomic/cmd
COPY ./internal /barcomic/internal
COPY ./scripts /barcomic/scripts

# Remove instances of robotgo from source
RUN sed -i '/robotgo/d' ./internal/barcomic/restapi.go

# Create new Go module without robotgo
RUN go mod init github.com/TheGrayDot/barcomic
RUN go get -v github.com/mdp/qrterminal@v1.0.1

# Compile application
RUN bash ./scripts/build_linux.sh

# Run application
CMD ./bin/barcomic-linux -v -a 0.0.0.0 -p 80 -s true -i false

# Hack to keep the container up and not complain
# CMD tail -f /dev/null
