
import uvicorn

from fastapi import APIRouter
from pydantic import BaseModel

from src.main import app, dummy_products, router as main_router
from src.main import Product

pact_router = APIRouter()


class ProviderState(BaseModel):
    state: str  # noqa: E999


@pact_router.post("/_pact/provider_states")
async def provider_states(provider_state: ProviderState):
    mapping = {
        "Product product0001 exists": setup_product0001,
    }
    mapping[provider_state.state]()

    return {"result": mapping[provider_state.state]}


# Make sure the app includes both routers. This needs to be done after the
# declaration of the provider_states
app.include_router(main_router)
app.include_router(pact_router)


def run_server():
    uvicorn.run(app)

def setup_product0001():
    productID = "product0001"
    price = 100.21
    stock = 100
    name = "Product 0001"

    dummy_products[productID] = Product(price=price, stock=stock)