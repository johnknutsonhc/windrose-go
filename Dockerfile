FROM golang:1.23-bookworm
# RUN go get -u github.com/jknutson/windrose-go
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV APP_HOME=/go/src/windrose-go
ARG APP_USER=app
ARG GROUP_ID=1001
ARG USER_ID=1001
# TODO: why does using $APP_USER not work below?
RUN groupadd --gid $GROUP_ID app && useradd -m -l --uid $USER_ID --gid $GROUP_ID app && \
      mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME
USER $APP_USER
WORKDIR $APP_HOME
COPY . .
RUN go build -buildvcs=false -o ./windrose-go .
EXPOSE 8080
HEALTHCHECK --interval=5m --timeout=3s \
  CMD curl -f http://localhost:8080/windrose || exit 1
CMD ["./windrose-go"]
