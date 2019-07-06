import React, { FC } from 'react';
import './style.scss';

interface ITagProps {

}

const Tag: FC<ITagProps> = ({ children }) => {
  return (
    <i className="tag">{children}</i>
  )
}

export default Tag;