
:: protoc.exe --go_out=plugins=grpc:./ service.proto
@echo off

echo build server message start
for %%f in (./proto/server/*.proto) do (
	echo build ./proto/server/%%f

	.\tool\protoc.exe --proto_path=./proto/server --proto_path=./proto/client --go_out=plugins=grpc:./ %%f
)
echo build server message finish

xcopy "./sparrow/protocol" "./" /s /i /y
rd /s/q sparrow

pause