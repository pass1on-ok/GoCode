FROM golang:latest

##buat folder APP
RUN mkdir /online-learning-platform

##set direktori utama
WORKDIR /online-learning-platform

##copy seluruh file ke completedep
ADD . /online-learning-platform

##buat executeable
RUN go build -o main .

##jalankan executeable
CMD ["./main"]
