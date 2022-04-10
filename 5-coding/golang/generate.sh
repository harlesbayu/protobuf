protoc --go_out=../complex --go_opt=paths=source_relative complex.proto
protoc --go_out=../enum_example --go_opt=paths=source_relative enum_example.proto
protoc --go_out=../simple --go_opt=paths=source_relative simple.proto