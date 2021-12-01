const { server } = require("./provider.js");

server.listen(8081, () => {
  console.log("Product Service listening on http://localhost:8081");
});
