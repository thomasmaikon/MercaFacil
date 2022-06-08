FROM golang

WORKDIR /API
RUN git clone https://github.com/thomasmaikon/MercaFacil.git
WORKDIR /API/MercaFacil/
RUN go mod download
CMD [ "go", "run" ,"main.go", "producao" ]