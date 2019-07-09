import React, { FC } from 'react';
import { Link } from 'react-router-dom';
import './style.scss';

import { IQuestion, IList } from '../../model';
import { difficultyOptions } from '../../constant';

export interface IQuestionsProps {
  data: IList<IQuestion>;
}

const Questions: FC<IQuestionsProps> = ({ data }) => {
  const { nodes } = data;

  return (
    <section className="questions">
      <table className="list">
        <thead>
          <tr>
            {/* <th>#</th> */}
            <th>题目</th>
            <th>难度</th>
          </tr>
        </thead>
        <tbody>
          {
            nodes.map(question => {
              return (
                <tr key={question.id}>
                  {/* <td>{question.id}</td> */}
                  <td className={question.summary && 'has-summary'}>
                    <Link to={`/question/${question.id}`}>
                      {question.title}
                      <br />
                      {question.summary && (
                        <span className="summary">{question.summary}</span>
                      )}
                    </Link></td>
                  <td>{difficultyOptions[question.difficulty - 1].text}</td>
                </tr>
              )
            })
          }
        </tbody>
      </table>

    </section>
  )
}

export default Questions;
