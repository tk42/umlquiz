// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'
import { API_URL, ENDPOINT_USER } from '../../components/const'

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
  const response = await fetch(API_URL+ENDPOINT_USER, {
    method: 'POST',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    body: req.body
  })
  const user = await response.json()
  res.status(200).json(user)
}
