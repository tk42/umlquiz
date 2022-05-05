// https://nextjs-ja-translation-docs.vercel.app/docs/api-routes/dynamic-api-routes
import type { NextApiRequest, NextApiResponse } from 'next'

export default async function handler(
    req: NextApiRequest,
    res: NextApiResponse<any[]>
  ) {
    const {
      query: { slug }
    } = req;
  
  };