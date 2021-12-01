class Repository {
  constructor() {
    this.entities = [];
  }

  fetchAll() {
    return this.entities;
  }

  getByProductId(productID) {
    return this.entities.find((entity) => productID == entity.productID);
  }

  insert(entity) {
    this.entities.push(entity);
  }

  clear() {
    this.entities = [];
  }
}

module.exports = Repository;
