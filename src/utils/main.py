# main.py

import os
import uvicorn
from fastapi import FastAPI
from auth_gateway.api import api
from auth_gateway.config import settings

def main():
    uvicorn.run(
        "auth_gateway.api:app",
        host=settings.HOST,
        port=settings.PORT,
        reload=True
    )

if __name__ == "__main__":
    main()