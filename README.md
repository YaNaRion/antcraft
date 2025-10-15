# antcraft

## TO GENERATE PROTO RUN AT PROJECT ROOT
- protoc -I=./proto --go_out=go-server/ ./proto/event_name.proto --cpp_out=cpp/proto

## TECH
- Client - Cpp with raylib
- Backend - Go
    - Communication with client via websocket and protobuff
- DB - Postgres
