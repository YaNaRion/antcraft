# antcraft

## TO GENERATE PROTO RUN AT PROJECT ROOT
- protoc -I=./proto --go_out=server/ ./proto/event_name.proto --cpp_out=client_cpp/proto

## TECH
- Client - Cpp with raylib
- Backend - Go
    - Communication with client via websocket and protobuff
- DB - Postgres
