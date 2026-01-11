# utils.py

import hashlib
import logging
import os
import secrets

from auth_gateway.config import Config

logger = logging.getLogger(__name__)

def generate_auth_code(length: int = 32) -> str:
    """Generate a random auth code."""
    return secrets.token_urlsafe(length)

def generate_token(length: int = 32) -> str:
    """Generate a random token."""
    return secrets.token_urlsafe(length)

def generate_secret_key(length: int = 32) -> str:
    """Generate a random secret key."""
    return secrets.token_urlsafe(length)

def generate_email_token(length: int = 32) -> str:
    """Generate a random email token."""
    return secrets.token_urlsafe(length)

def hash_password(password: str) -> str:
    """Hash a password using SHA256."""
    return hashlib.sha256(password.encode()).hexdigest()

def get_file_contents(file_path: str) -> str:
    """Read the contents of a file."""
    return open(file_path, "r").read()

def get_file_mode(file_path: str) -> str:
    """Get the file mode of a file."""
    return os.stat(file_path).st_mode