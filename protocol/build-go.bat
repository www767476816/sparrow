

:: protoc.exe --go_out=./ service.proto
@echo off

echo build go start
for %%f in (./proto/go/*.proto) do (
	echo build ./proto/go/%%f.proto
	.\tool\protoc.exe -I=./proto/go --go_out=go_protocol %%f
)
echo build go finish

pause