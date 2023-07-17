FROM golang AS build
WORKDIR /home
COPY . app
RUN ls -la
WORKDIR /home/app
RUN ls -la
RUN go mod tidy
RUN go build -o /home/app/main /home/app/cmd/main.go

FROM alpine:latest AS deploy
WORKDIR /home
RUN mkdir app
COPY ./run.sh /home/app/run.sh
COPY --from=build /home/app/main /home/app
WORKDIR /home/app
RUN ls -la
ENTRYPOINT [ "/bin/sh", "/home/app/run.sh" ]
EXPOSE 8080