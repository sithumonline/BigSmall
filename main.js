const http = require("http");

const server = http.createServer((req, res) => {
  if (req.url === "/small") {
    console.log("got /small request");
    res.end("Small is beautiful!");
  }
  if (req.url === "/big") {
    // print 1 to 50000
    for (let i = 0; i < 50000; i++) {
      console.log(`${i}, `);
    }
    console.log("got /big request");
    res.end("Big is beautiful!");
  }
});

server.listen(3000);

// wrk --rate=20 --duration=5s "http://localhost:3000/big" & wrk --rate=20 --duration=5s "http://localhost:3000/small"
