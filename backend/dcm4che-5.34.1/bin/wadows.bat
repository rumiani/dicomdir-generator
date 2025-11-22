@echo off
rem -------------------------------------------------------------------------
rem wadows  Launcher
rem -------------------------------------------------------------------------

if not "%ECHO%" == ""  echo %ECHO%
if "%OS%" == "Windows_NT"  setlocal

set MAIN_CLASS=org.dcm4che3.tool.wadows.WadoWS
set MAIN_JAR=dcm4che-tool-wadows-5.34.1.jar

set DIRNAME=.\
if "%OS%" == "Windows_NT" set DIRNAME=%~dp0%

rem Read all command line arguments

set ARGS=
:loop
if [%1] == [] goto end
        set ARGS=%ARGS% %1
        shift
        goto loop
:end

if not "%DCM4CHE_HOME%" == "" goto HAVE_DCM4CHE_HOME

set DCM4CHE_HOME=%DIRNAME%..

:HAVE_DCM4CHE_HOME

if not "%JAVA_HOME%" == "" goto HAVE_JAVA_HOME

set JAVA=java

goto SKIP_SET_JAVA_HOME

:HAVE_JAVA_HOME

set JAVA=%JAVA_HOME%\bin\java

:SKIP_SET_JAVA_HOME

set CP=%DCM4CHE_HOME%\etc\wadows\
set CP=%CP%;%DCM4CHE_HOME%\lib\%MAIN_JAR%
set CP=%CP%;%DCM4CHE_HOME%\lib\dcm4che-core-5.34.1.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\dcm4che-mime-5.34.1.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\dcm4che-tool-common-5.34.1.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\dcm4che-xdsi-5.34.1.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\slf4j-api-2.0.16.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\logback-core-1.5.12.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\logback-classic-1.5.12.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\commons-cli-1.9.0.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\jakarta.activation-api-2.1.3.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\angus-activation-2.0.2.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\jakarta.xml.bind-api-4.0.2.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\jaxb-runtime-4.0.5.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\jaxb-core-4.0.5.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\jakarta.xml.ws-api-4.0.2.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\rt-4.0.2.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\jakarta.xml.soap-api-3.0.1.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\istack-commons-runtime-4.1.2.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\saaj-impl-3.0.3.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\streambuffer-2.1.0.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\policy-4.0.2.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\gmbal-api-only-4.0.3.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\mimepull-1.10.0.jar
set CP=%CP%;%DCM4CHE_HOME%\lib\stax-ex-2.1.0.jar

"%JAVA%" %JAVA_OPTS% -cp "%CP%" %MAIN_CLASS% %ARGS%
