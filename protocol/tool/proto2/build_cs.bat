SET name=CenterServer

::编译
protogen -i:%name%.proto -o:%name%.cs  
::::::::::::::::::::::::::::::::::::::::::::::::::::::
::拷贝
XCOPY %name%.cs ..\message_cs\ /Y
DEL %name%.cs
::::::::::::::::::::::::::::::::::::::::::::::::::::::
PAUSE