import React, { FC } from 'react';
import { Link } from 'react-router-dom';
import './style.scss';

import { IQuestion, IList } from '../../model';

export interface IQuestionsProps {
  data: IList<IQuestion>;
}

const Questions: FC<IQuestionsProps> = ({ data }) => {

  const { totalCount, nodes } = data;

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
                  <td><Link to={`/question/${question.id}`}>{question.title}</Link></td>
                  <td>{question.difficulty}</td>
                </tr>
              )
            })
          }
        </tbody>
      </table>
      <p>共{totalCount}个题目</p>
    </section>
  )
}

export default Questions;
