FROM busybox

MAINTAINER gzh guozh318@163.com
COPY todo_backend ./todo_backend

EXPOSE 80

ENTRYPOINT ["./todo_backend"]