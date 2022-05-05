export type FetchQuizProps = {
    access_token: string
    language: string
    quiz_id?: string
    user_id?: string
    linkto: string
}

export type User = {
    user_id: string
    username: string
    password: string
    email: string
    profile: string
    created_at: number
    updated_at: number
    membership: string
    liked_quiz_ids: string
    quiz_history: string
}

export type Quiz = {
    quiz_id: string
    language: string
    diagram_type: string
    level: string
    title: string
    text: string
    diagram: string
    likes: number
    author_id: string
    created_at: number
    updated_at: number
    status: string
}