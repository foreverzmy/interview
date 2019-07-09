import React, { FC } from 'react';
import { Link } from 'react-router-dom';
import './style.scss';

import { IQuestion, IList } from '../../model';
import { difficultyOptions } from '../../constant';

import Pagination from '../Pagination'

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
                  <td>{difficultyOptions[question.difficulty - 1].text}</td>
                </tr>
              )
            })
          }
        </tbody>
      </table>
      <Pagination
        current={1}
        size={20}
        total={totalCount}
        onChange={(current) => console.log(current)}
      />
    </section>
  )
}

export default Questions;
