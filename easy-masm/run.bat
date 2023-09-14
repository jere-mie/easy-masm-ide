@echo off
cd easy-masm
copy src\%1.asm lib\
cd lib
echo on
aml.exe /c /Zd /coff %1.asm
alink.exe /SUBSYSTEM:CONSOLE %1.obj
%1.exe
@echo off
del %1.exe
del %1.obj
del %1.asm
cd ..