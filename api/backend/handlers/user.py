import uuid
import time
from fastapi import Form

from . import session_scope, User


def post(
        username: str,
        password: str,
        email: str
    ):
    now = int(time.time())
    user = User(
        username=username,
        password=password,
        email=email,
        user_id=str(uuid.uuid4()),
        created_at=now,
        updated_at=now,
        membership='bronze',
        liked_quiz_ids='',
        quiz_history='',
    )
    with session_scope() as s:
        s.add(user)
        return user.to_dict()

def get(
        user_id: str,
    ):
    with session_scope() as s:
        user = s.query(User).filter(User.user_id==user_id).first()
        return user.to_dict()

def put(
        user_id: str,
        user: dict,
    ):
    with session_scope() as s:
        u = s.query(User).filter(User.user_id==user_id).first()
        for k, v in user.items():
            setattr(u, k, v)
    return user

def delete(
        user_id: str,
    ):
    with session_scope() as s:
        user = s.query(User).filter(User.user_id==user_id).first()
        s.delete(user)
        return user.to_dict()
