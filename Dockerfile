#=== Build FrontEnd ===#
FROM node:14-alpine AS frontend-build
WORKDIR /app
COPY ./frontend/package*.json ./
RUN npm install --only=production
COPY ./frontend .
RUN npm run build

#=== Build BackEnd ===#
FROM golang:1.17-alpine AS go-build
WORKDIR /app
COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod download
COPY ./backend .
RUN go build -o main .

#=== Build Image ===#
FROM alpine
WORKDIR /app
COPY --from=frontend-build /app/build ./web
COPY --from=go-build /app/main  ./
COPY --from=go-build /app/conf/config.json ./conf/config.json
CMD ["./main"]