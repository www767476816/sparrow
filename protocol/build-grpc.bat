
:: protoc.exe --go_out=plugins=grpc:./ service.proto
@echo off

echo build grpc start
for %%f in (./proto/grpc/*.proto) do (
	echo build ./proto/grpc/%%f

	.\tool\protoc.exe -I=./proto/grpc -I=./proto/go --go_out=plugins=grpc:./ %%f
)
echo build grpc finish

pause