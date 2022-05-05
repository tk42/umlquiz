// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'

type Todo = {
  userId: number
  id: number
  title: string
  completed: boolean
}

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<Todo[]>
) {
  const response = await fetch('https://jsonplaceholder.typicode.com/todos/')
  const todos = await response.json()
  res.status(200).json(todos)
}
