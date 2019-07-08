import React, { FC } from 'react';
import Coder from '../Coder';
import Markdown from '../Markdown';
import './style.scss';

interface IEditorProps {
  value: string;
  onChange: (value: string) => void;
}

const Editor: FC<IEditorProps> = ({ value, onChange }) => {

  return (
    <article className="editor">
      <Coder value={value} onChange={onChange} />
      <Markdown text={value} />
    </article>
  )
}

export default Editor;
