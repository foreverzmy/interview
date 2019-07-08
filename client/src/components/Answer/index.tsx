import React, { FC } from 'react';
import './style.scss';

import { IAnawer } from '../../model';

export interface IAnswerProps {
  data: IAnawer;
}

const Answer: FC<IAnswerProps> = ({ data }) => {
  return (
    <section>
      {data.content}
    </section>
  )
}

export default Answer;
