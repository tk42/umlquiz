import time
from .oauth2 import create_access_token
from . import session_scope, User

def get(
    username: str,
    password: str,
):
    with session_scope() as s:
        user = s.query(User).filter(User.username==username and User.password==password).first()
        if user is None:
            raise HTTPException(
                status_code=status.HTTP_404_NOT_FOUND,
                detail='Failed Authentication'
            )
        access_token = create_access_token(data={'sub': username})
        return {
            'access_token': access_token,
            'token_type': 'bearer',
        }
