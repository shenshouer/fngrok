FROM scratch

COPY ./bin/server/fngrok /fngrok

EXPOSE 8001

ENTRYPOINT ["/fngrok"]