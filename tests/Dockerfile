FROM golangci/golangci-lint:v1.31-alpine


ARG NAME_ENV=report-creator

RUN mkdir -p /${NAME_ENV}

ADD go.mod /${NAME_ENV}
ADD go.sum /${NAME_ENV}
ADD ./docs /${NAME_ENV}/docs
RUN cd /${NAME_ENV} && go mod download


ADD ./app /${NAME_ENV}/app
WORKDIR /${NAME_ENV}

RUN golangci-lint run --issues-exit-code=1 --deadline=600s ./...

RUN go test -cover ./...
