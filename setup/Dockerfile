FROM golang
WORKDIR /home
COPY . app
RUN ls 
WORKDIR /home/app
RUN ls 
RUN go mod tidy
RUN go build -o /home/app/main /home/app/cmd/main.go
ENTRYPOINT [ "/home/app/run.sh" ]
EXPOSE 8080