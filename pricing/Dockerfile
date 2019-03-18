FROM iron/base

EXPOSE 50051
ADD pricing-linux-amd64 /

ARG serverIp
ENV serverIpArg=-serverIp=$serverIp

ENTRYPOINT ./pricing-linux-amd64 $serverIpArg
