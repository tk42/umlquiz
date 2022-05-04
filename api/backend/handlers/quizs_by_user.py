from fastapi import Form
from . import session_scope, Quiz, User


def get(
        user_id: str
    ):
    with session_scope() as s:
        quizs = s.query(User).filter(User.user_id==user_id).all()
        return [quiz.to_dict() for quiz in quizs]
