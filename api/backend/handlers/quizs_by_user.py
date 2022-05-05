from fastapi import Form
from . import session_scope, Quiz, User


def get(
        user_id: str
    ):
    with session_scope() as s:
        quizs = s.query(Quiz).filter(Quiz.author_id==user_id).all()
        return [quiz.to_dict() for quiz in quizs]
