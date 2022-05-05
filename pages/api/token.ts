import { API_URL, ENDPOINT_TOKEN } from '../../components/const';
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
    const response = await fetch(API_URL+ENDPOINT_TOKEN, {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: "username="+process.env.AUTH_USERNAME+"&password="+process.env.AUTH_PASSWORD,
    })
    const data = await response.json()
    res.status(200).json(data)
}
