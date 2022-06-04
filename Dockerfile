FROM golang

WORKDIR /API
COPY . .

RUN go mod download
RUN cd tests/
CMD [ "go", "test" ]