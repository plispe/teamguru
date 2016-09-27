FROM alpine:3.4
MAINTAINER Petr Pliska <petr.pliska@be.plus>
RUN apk add --update bash
COPY artifact /
ENV PORT 80
EXPOSE 80
CMD ["./artifact"]
