import React, { useEffect, useState } from 'react';
const App = () => {
  const [count, setCount] = useState(0);
  const [questions, setQuestions] = useState([] as any[]);

  useEffect(() => {
    fetch('/graphql', {
      method: 'POST',
      mode: 'cors',
      headers: {
        ContentType: 'application/json',
      },
      body: JSON.stringify({
        query: `
          {
            questions(page: 1, size: 10, keyword: "") {
              totalCount
              nodes{
                id
                title
                summary
                content
                difficulty
              }
            }
          }
        `
      })
    })
      .then((res) => res.json())
      .then(({ data }) => {
        const { questions } = data;
        setCount(questions.totalCount)
        setQuestions(questions.nodes)
      })
  }, [])

  return (
    <article>
      <p>{count}</p>
      <ul>
        {
          questions.map(qs => (
            <details key={qs.id}>
              <summary>{qs.title}</summary>
              <p>{qs.content}</p>
            </details>
          ))
        }
      </ul>
    </article>
  )
}

export default App;