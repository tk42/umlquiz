from fastapi import Form
from . import session_scope, Quiz, QuizText


def get(
        language: str,
        user_id: str
    ):
    with session_scope() as s:
        quizs = s.query(Quiz).filter(Quiz.author_id==user_id).all()
        quiztexts = s.query(QuizText).filter(QuizText.language==language).all()
        mq = {q.quiz_id: q.to_dict() for q in quizs}
        return [qt.to_dict() | mq[qt.quiz_id] for qt in quiztexts]
