import difflib

def diff(correct: str, answer: str):
    res = difflib.ndiff(correct.split(), answer.split())
    return '\n'.join(res)
