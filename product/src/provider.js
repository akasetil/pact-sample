const express = require("express");
const cors = require("cors");
const bodyParser = require("body-parser");
const Repository = require("./productRepository");

const server = express();
server.use(cors());
server.use(bodyParser.json());
server.use(
  bodyParser.urlencoded({
    extended: true,
  })
);
server.use((req, res, next) => {
  res.header("Content-Type", "application/json; charset=utf-8");
  next();
});

const productRepository = new Repository();

// setup用の関数
const createProduct0001 = () => {
  const data = {
    productID: "product0001",
    price: 1000,
    stock: 20,
  };

  productRepository.insert(data);
};

// 検証用のAPI
server.get("/product/:productid", (req, res) => {
  const response = productRepository.getByProductId(req.params.productid);
  if (response) {
    res.end(JSON.stringify(response));
  } else {
    res.writeHead(404);
    res.end();
  }
});

module.exports = {
  server,
  createProduct0001,
  productRepository,
};
