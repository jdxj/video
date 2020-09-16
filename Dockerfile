FROM alpine
ADD abc /abc
ENTRYPOINT [ "/abc" ]
