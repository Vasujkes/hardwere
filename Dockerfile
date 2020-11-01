FROM golang
WORKDIR /tutorial/GoRest
ENV GOPATH=/go
COPY . /tutorial/GoRest
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gorilla/handlers
RUN go build -o main .
CMD [ "./main" ]