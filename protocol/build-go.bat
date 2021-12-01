

:: protoc.exe --go_out=./ service.proto
@echo off

echo build go start
for %%f in (./proto/go/*.proto) do (
	echo build ./proto/go/%%f

	.\tool\protoc.exe -I=./proto/go --go_out=./ %%f
)
echo build go finish

pause