from fastapi import Form
from . import session_scope, ReportQuiz
from diff import diff

def post(
        quiz_id: str,
        language: str,
        answer_diagram: str,
    ):
    with session_scope() as s:
        correct_quiz = s.query(Quiz).filter(Quiz.quiz_id==quiz_id and Quiz.language==language).first()
        return diff(correct_quiz.diagram, answer_diagram)

def get(
    ):
    with session_scope() as s:
        quizs = s.query(Quiz).all()
        return [quiz.to_dict() for quiz in quizs]

def put(
        quiz_id: str,
        language: str,
        report_test: str,
        report_diagram: str,
    ):
    now = int(time.time())
    q = ReportQuiz(
        quiz_id = quiz_id,
        language = language,
        text = report_test,
        diagram = report_diagram,
        created_at = now,
    )
    with session_scope() as s:
        s.add(q)
        return q.to_dict()
