

:: protoc.exe --go_out=./ service.proto
@echo off

echo build client message start
for %%f in (./proto/client/*.proto) do (
	echo build ./proto/client/%%f

	.\tool\protoc.exe --proto_path=./proto/client/ --go_out=./ %%f
)
echo build client message finish

xcopy "./sparrow/protocol" "./" /s /i /y
rd /s/q sparrow
pause