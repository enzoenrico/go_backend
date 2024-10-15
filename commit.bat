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

REM Prompt the user for a commit message or use the last edited file as default
SET /P "user_commit_message=Enter a commit message (leave blank for default): "

REM Determine the commit message

IF "%user_commit_message%"=="" (
    SET "commit_message=Updated: %changed_files%"
) ELSE (
    SET "commit_message=%user_commit_message%"
)

REM Commit the changes
git commit -m "%commit_message%" --allow-empty

REM Push the changes to the master branch
git push origin master
