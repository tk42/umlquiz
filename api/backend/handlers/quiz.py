from fastapi import Form
from . import session_scope, ReportQuiz
from diff import diff

# TODO: remove diff. just post id and add this to user history
def post(
        quiz_id: str,
        language: str,
        answer_diagram: str,
    ):
    with session_scope() as s:
        correct_quiz = s.query(Quiz).filter(Quiz.quiz_id==quiz_id and Quiz.language==language).first()
        return diff(correct_quiz.diagram, answer_diagram)

def get(
        language: str,
    ):
    with session_scope() as s:
        quizs = s.query(Quiz).all()
        return [quiz.to_dict() for quiz in quizs if quiz['language'] == language]

def put(
        quiz_id: str,
        language: str,
        report_text: str,
        report_diagram: str,
    ):
    now = int(time.time())
    q = ReportQuiz(
        quiz_id = quiz_id,
        language = language,
        text = report_text,
        diagram = report_diagram,
        created_at = now,
    )
    with session_scope() as s:
        s.add(q)
        return q.to_dict()
