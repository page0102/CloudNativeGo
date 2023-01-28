FROM ubuntu
RUN echo '这是一个基于ubuntu且应用为httpserver的镜像' && mkdir -p /dockerstudy/httpserver
COPY go/gohttpserver /dockerstudy/httpserver/
#此处source路径必须是相对路径，使用绝对路径会报错COPY failed: file not found in build context or excluded by .dockerignore: stat /home/zhiyong/dockerstudy/httpserver/go/gohttpserver: file does not exist

ENTRYPOINT ["/dockerstudy/httpserver/gohttpserver"]
