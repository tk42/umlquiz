from fastapi import Form
from . import session_scope, Quiz
from diff import diff

def post(
        quiz_id: str,
        language: str,
        answer_quiz: dict,
    ):
    with session_scope() as s:
        correct_quiz = s.query(Quiz).filter(Quiz.quiz_id==quiz_id and Quiz.language==language).first()
        return diff(correct_quiz['diagram'], answer_quiz['diagram'])

def get(
    ):
    with session_scope() as s:
        quizs = s.query(Quiz).all()
        return [quiz.to_dict() for quiz in quizs]

def put(
        quiz_id: str,
        quiz: dict,
    ):
    with session_scope() as s:
        q = s.query(Quiz).filter(Quiz.quiz_id==quiz_id).first()
        for k, v in quiz.items():
            setattr(q, k, v)
    return quiz
