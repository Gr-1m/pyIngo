@echo off

for /F "delims=" %i in ('cd') do set pwd=%i
for /F "delims=" %i in ('go env GOROOT') do set gosrc=%i\src

echo %pwd
echo %gosrc

Rem mklink 
mklink /J "%gosrc%\pyIngo" %pwd%
