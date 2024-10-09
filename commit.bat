@echo off

REM Stage all changes
git add .

REM Capture the list of changed files
SET "changed_files="
FOR /F "delims=" %%f IN ('git diff --cached --name-only') DO (
    SET "changed_files= %%f"
)

REM Check if there are any changed files
IF "%changed_files%"=="" (
    echo No changes to commit.
    exit /b 0
)

REM Create a custom commit message
SET "commit_message=Changed files: %changed_files%"

REM Commit the changes
git commit -m "%commit_message%"

REM Push the changes to the master branch
git push origin master
