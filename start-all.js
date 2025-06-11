const concurrently = require("concurrently");

concurrently([
  { command: "cd frontend && npm run dev", name: "frontend", prefixColor: "blue" },
  { command: "cd backend && go run main.go", name: "backend", prefixColor: "green" }
]);