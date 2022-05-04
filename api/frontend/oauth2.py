# https://self-methods.com/fastapi-authentication/
from jose import jwt
from datetime import timedelta

from fastapi.security import OAuth2PasswordBearer
from typing import Optional
from datetime import datetime, timedelta
  
# openssl rand -hex 32
SECRET_KEY = "3d6a9e1c95d1f92a51d674e0f6589d6a568bf43b2d3463d7551ff0b93b78264c"
ALGORITHM = "HS256"

def create_access_token(data: dict, expires_delta: Optional[timedelta] = None):
    to_encode = data.copy()
    if expires_delta:
        expire = datetime.utcnow() + expires_delta
    else:
        expire = datetime.utcnow() + timedelta(minutes=60)
    to_encode.update({"exp": expire})
    encoded_jwt = jwt.encode(to_encode, SECRET_KEY, algorithm=ALGORITHM)
    return encoded_jwt
