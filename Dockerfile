FROM golang:1.16.5-buster
# RUN go get -u github.com/jknutson/windrose-go
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV APP_USER app
ENV APP_HOME /go/src/windrose-go
ARG GROUP_ID=1001
ARG USER_ID=1001
RUN groupadd --gid $GROUP_ID && useradd -m -l --uid $USER_ID --gid $GROUP_ID $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME
USER $APP_USER
WORKDIR $APP_HOME
COPY . .
RUN go build -o ./windrose-go .
EXPOSE 8080
CMD ["./windrose-go"]
