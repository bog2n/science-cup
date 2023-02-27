// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'

type Data = {
  name: string
}

export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
	console.log("ASDS")
  res.status(200).json(
  {
  "ok": true,
  "proteins": [
    {
      "protein": "MGCLCTRETASYCLSNATQFFLRIDHDLHTFELESIMTSYYYCVLTSLNSRVMLI",
      "mass": 6402.010031085999,
      "hindex": 20.699999999999996,
      "isopoint": 6.9453125,
      "ph": 2.964624558967162,
      "polarity": -1
    },
    {
      "protein": "MKSAAGRRKAFSTCGSHLTVVSVFYGTLIFMYVQPKYSQTFETDKVSSIFYTLVIPMLNPLIYSLRNKDAKDAIERTWKKIVTSFS",
      "mass": 9818.137355886,
      "hindex": 38.88999999999999,
      "isopoint": 10.8828125,
      "ph": 2.9552879568164445,
      "polarity": 11
    },
    {
      "protein": "MTIHFYCKPTLVSS",
      "mass": 1625.7945706859998,
      "hindex": 8.650000000000002,
      "isopoint": 8.1484375,
      "ph": 2.2999913301215726,
      "polarity": 1
    },
    {
      "protein": "MCVIEFYDVWESQRDFELFV",
      "mass": 2554.1548688860003,
      "hindex": 16.74,
      "isopoint": 3.9921875,
      "ph": 2.8060362962946956,
      "polarity": -3
    },
    {
      "protein": "MCCLQMKFDYDTKLLISVHYLRVFY",
      "mass": 3128.535983286,
      "hindex": 12.43,
      "isopoint": 9.2421875,
      "ph": 2.5295663794597942,
      "polarity": 3
    },
  ],
  }
  )
}
