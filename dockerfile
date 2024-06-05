FROM golang

WORKDIR /usr/src/app

COPY go.mod ./ 
RUN go mod download && go mod verify

COPY . . 

RUN go build -o app main.go

CMD [ "/usr/src/app/app", "server" ]


