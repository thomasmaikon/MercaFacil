FROM golang

WORKDIR /API
COPY . .

RUN go mod download
CMD [ "go", "test" ,"/API/tests/" ]