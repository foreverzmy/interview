import React, { FC, useCallback, ChangeEvent, KeyboardEvent } from 'react';
import './style.scss';

interface ISearchProps {
  placeholder: string;
  value: string;
  onChange: (value: string) => void;
  onEnter?: (value: string) => void;
}

// 输入框组件
const Search: FC<ISearchProps> = ({ placeholder, value, onChange, onEnter }) => {
  const handleChange = useCallback((evt: ChangeEvent<HTMLInputElement>) => {
    onChange(evt.target.value);
  }, [])

  // enter 事件
  const handleKeyDown = useCallback((evt: KeyboardEvent<HTMLInputElement>) => {
    if (evt.keyCode === 13) {
      onEnter && onEnter(value);
    }
  }, [value])

  return (
    <section className="search">
      <input
        className="search__input"
        placeholder={placeholder}
        value={value}
        onChange={handleChange}
        onKeyDown={handleKeyDown}
      />
    </section>
  )
}

export default Search;