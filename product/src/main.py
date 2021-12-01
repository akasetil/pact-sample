from fastapi import FastAPI, HTTPException, APIRouter

class Product:
    def __init__(self, productID=None, price=0, stock=0, name=None):
        self.productID = productID
        self.price = price
        self.stock = stock
        self.Name = name

    def reduceStock(self, quantity):
        self.stock -= quantity


router = APIRouter()
app = FastAPI()

dummy_products = {
    "aaa": Product(productID="aaa", price=1000, stock=100),
    "bbb": Product(productID="bbb", price=2000, stock=100),
    "ccc": Product(productID="ccc", price=3000, stock=100),
}


@app.get("/product/{product_id}")
async def get_product(product_id: str):
    if product_id not in dummy_products:
        raise HTTPException(status_code=404, detail="product not found")
    return dummy_products[product_id]
