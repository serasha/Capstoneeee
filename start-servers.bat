@echo off
echo Starting Backend Server...
cd "c:\Data D\Capstone Project\Capstoneeee\backend"
start "Backend Server" cmd /k "go run main.go"

timeout /t 3

echo Starting Frontend Server...
cd "c:\Data D\Capstone Project\Capstoneeee\frontend"
start "Frontend Server" cmd /k "npm run dev"

echo Both servers are starting...
echo Backend: http://localhost:8070
echo Frontend: http://localhost:5173
pause
