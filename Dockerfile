FROM golang

WORKDIR /API
RUN git clone https://github.com/thomasmaikon/MercaFacil
WORKDIR /API/MercaFacil/
RUN go mod download
EXPOSE 8000
CMD [ "go", "run" ,"main.go" ]