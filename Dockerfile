FROM alpine:3.15.0
# RUN apk add --no-cache ca-certificate=20211220-r0
RUN apk add --no-cache tzdata

ENV TZ=Asia/Bangkok
ENV HOST=0.0.0.0

# Copy binary to image
COPY /server /server

CMD ["/server"]

