FROM golang

WORKDIR /API
COPY . .
EXPOSE 8000
CMD [ "go", "run" ,"main.go", "producao" ]