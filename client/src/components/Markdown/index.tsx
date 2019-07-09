import React, { FC, memo, useMemo } from 'react';
import marked from 'marked';
import DOMPurify from 'dompurify';
import './style.scss';

export interface IMarkdownProps {
  text: string;
}

const Markdown: FC<IMarkdownProps> = ({ text }) => {
  const rawMarkup = useMemo(() => {
    return {
      __html: marked(DOMPurify.sanitize(text)),
    }
  }, [text])

  return (
    <div className="markdown-wrapper" dangerouslySetInnerHTML={rawMarkup} />
  )
}

export default memo(Markdown);
